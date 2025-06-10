<template>
  <q-page padding class="q-gutter-md">
    <!-- Top Row -->
    <q-card class="full-width">
      <q-card-section>
        <div class="text-h6">Defined McpServers</div>
        <div
          v-for="server in servers"
          :key="server.id"
          class="row items-center q-mb-sm"
        >
          <div class="q-mr-md">{{ server.name }}</div>
          <q-badge color="primary" :label="`Prompts: ${server.prompt_count}`" />
          <q-badge color="secondary" :label="`Resources: ${server.resource_count}`" />
          <q-badge color="deep-orange" :label="`Tools: ${server.tool_count}`" />

          <q-btn flat round icon="play_arrow" class="q-ml-sm" size="sm" />
          <q-btn flat round icon="stop" class="q-ml-sm" size="sm" />
          <q-btn flat round icon="info" class="q-ml-sm" size="sm" />
        </div>
      </q-card-section>
    </q-card>

    <!-- Middle Row: Icons -->
    <div class="row q-gutter-sm q-my-md">
      <q-btn flat icon="memory" label="New McpServer" draggable="true" @dragstart="startDrag('server')" />
      <q-btn flat icon="description" label="Prompt" draggable="true" @dragstart="startDrag('prompt')" />
      <q-btn flat icon="inventory_2" label="Resource" draggable="true" @dragstart="startDrag('resource')" />
      <q-btn flat icon="build_circle" label="Tool" draggable="true" @dragstart="startDrag('tool')" />
    </div>

    <!-- Bottom Row: SVG Canvas -->
    <div
      ref="canvasContainer"
      class="canvas-container bg-grey-2"
      @dragover.prevent
      @drop="handleDrop"
    >
      <svg ref="svgCanvas" width="100%" height="600px"></svg>
    </div>
  </q-page>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as d3 from 'd3'

const servers = ref([])
const svgCanvas = ref(null)
const canvasContainer = ref(null)
const draggedType = ref(null)

function startDrag(type) {
  draggedType.value = type
}

function handleDrop(event) {
  const coords = d3.pointer(event, svgCanvas.value)
  if (draggedType.value === 'server') {
    if (svgCanvas.value && svgCanvas.value.children.length > 0) {
      const shouldClear = confirm('Canvas is not empty. Clear existing content?')
      if (!shouldClear) return
      svgCanvas.value.innerHTML = ''
    }
    promptServerInfo(coords)
  } else {
    createNode(draggedType.value, coords)
  }
}

function promptServerInfo([x, y]) {
  const id = prompt('Enter server ID:')
  const addr = prompt('Enter address:')
  const version = prompt('Enter version:')
  const transport = prompt('Enter transport (http, stdio, in-memory):')
  const buildType = prompt('Enter build type (binary, container):')
  if (!id || !addr || !version || !transport || !buildType) return

  const serverNode = { id, addr, version, transport, buildType }
  console.debug('[ServerBuilder] Created server:', serverNode)

  const svg = d3.select(svgCanvas.value)
  const group = svg.append('g').attr('transform', `translate(${x},${y})`)

  group.append('rect')
    .attr('width', 100)
    .attr('height', 50)
    .attr('fill', '#ab47bc')

  group.append('text')
    .attr('x', 50)
    .attr('y', 25)
    .attr('text-anchor', 'middle')
    .attr('alignment-baseline', 'middle')
    .text(serverNode.id)
    .style('fill', '#fff')
}

function createNode(type, [x, y]) {
  const svg = d3.select(svgCanvas.value)
  const group = svg.append('g').attr('transform', `translate(${x},${y})`)

  group.append('circle')
    .attr('r', 30)
    .attr('fill', typeColor(type))

  group.append('text')
    .attr('text-anchor', 'middle')
    .attr('y', 5)
    .text(type.toUpperCase())
    .style('font-size', '10px')
    .style('fill', '#fff')
}

function typeColor(type) {
  return {
    prompt: '#42a5f5',
    resource: '#66bb6a',
    tool: '#ffa726',
    server: '#ab47bc'
  }[type] || '#90a4ae'
}

function updateList(newList) {
  servers.value = newList
  console.debug('[ServerBuilder] Updating server list:', newList)
}

onMounted(() => {
  const svg = d3.select(svgCanvas.value)
  svg.append('rect')
    .attr('width', '100%')
    .attr('height', '100%')
    .attr('fill', 'transparent')

  if (!window.__refs) window.__refs = {}

  window.__refs = {
    ...window.__refs,
    gauge_client: { value: null },
    gauge_server: { value: null },
    gauge_discovery: { value: null },
    list_client: { value: null },
    list_discovery: { value: null },
    list_server: { value: { updateList } },
    console: { value: null },
  }

  console.debug('[ServerBuilder] Registered refs into window.__refs')

  if (Array.isArray(window.__sseBuffer)) {
    const flush = window.__sseBuffer.filter(msg => msg.msg_id === 'list_server_ops')
    flush.forEach(({ payload }) => updateList(payload))
    window.__sseBuffer = window.__sseBuffer.filter(msg => msg.msg_id !== 'list_server_ops')
    if (flush.length > 0) {
      console.debug(`[ServerBuilder] Flushed ${flush.length} buffered list_server_ops messages`)
    }
  }
})
</script>

<style scoped>
.canvas-container {
  border: 1px solid #ccc;
  min-height: 600px;
  position: relative;
}
</style>
