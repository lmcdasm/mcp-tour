<template>
  <q-page padding class="q-gutter-md">
    <!-- Top Row: List of Defined Servers -->
    <q-card class="full-width">
      <q-card-section>
        <div class="text-h6">Mcp Servers</div>
        <div
          v-for="server in servers"
          :key="server.id"
          class="row items-center q-mb-sm justify-between"
        >
          <div class="col-auto text-subtitle2 q-mr-md">{{ server.name }}</div>
          <q-badge color="primary" :label="`Prompts: ${server.prompt_count ?? 0}`" />
          <q-badge color="secondary" :label="`Resources: ${server.resource_count ?? 0}`" />
          <q-badge color="deep-orange" :label="`Tools: ${server.tool_count ?? 0}`" />

          <div class="q-gutter-xs">
            <q-btn dense flat icon="play_arrow" @click="start(server)" />
            <q-btn dense flat icon="stop" @click="stop(server)" />
            <q-btn dense flat icon="info" @click="loadToCanvas(server)" />
          </div>
        </div>
      </q-card-section>
    </q-card>

    <!-- Middle Row: Draggable Tool Buttons -->
    <div class="row q-gutter-sm q-my-md">
      <q-btn flat icon="memory" label="New McpServer" @click="spawnServer" />
      <q-btn flat icon="description" label="Prompt" draggable @dragstart="startDrag('prompt')" />
      <q-btn flat icon="inventory_2" label="Resource" draggable @dragstart="startDrag('resource')" />
      <q-btn flat icon="build_circle" label="Tool" draggable @dragstart="startDrag('tool')" />
    </div>

    <!-- Bottom Row: Interactive SVG Canvas -->
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

function updateList(newList) {
  servers.value = newList
  console.debug('[ServerBuilder] Updating server list:', newList)
}

function spawnServer() {
  console.debug('[ServerBuilder] New McpServer creation requested.')
}

function start(server) {
  console.debug('[ServerBuilder] Start server:', server)
}

function stop(server) {
  console.debug('[ServerBuilder] Stop server:', server)
}

function loadToCanvas(server) {
  console.debug('[ServerBuilder] Load server to canvas:', server)
}

const svgCanvas = ref(null)
const canvasContainer = ref(null)
const draggedType = ref(null)

function startDrag(type) {
  draggedType.value = type
}

function handleDrop(event) {
  const coords = d3.pointer(event, svgCanvas.value)
  createNode(draggedType.value, coords)
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

onMounted(() => {
  const svg = d3.select(svgCanvas.value)
  svg.append('rect')
    .attr('width', '100%')
    .attr('height', '100%')
    .attr('fill', 'transparent')

  // Init __refs safely
  if (!window.__refs) window.__refs = {}

  // Register list_server with live handler
  window.__refs = {
    ...window.__refs,
    gauge_client: window.__refs.gauge_client ?? { value: null },
    gauge_server: window.__refs.gauge_server ?? { value: null },
    gauge_discovery: window.__refs.gauge_discovery ?? { value: null },
    list_client: window.__refs.list_client ?? { value: null },
    list_discovery: window.__refs.list_discovery ?? { value: null },
    list_server: { value: { updateList } },
    console: window.__refs.console ?? { value: null }
  }

  console.debug('[ServerBuilder] Registered refs into window.__refs')

  // Flush buffered list_server_ops messages
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

