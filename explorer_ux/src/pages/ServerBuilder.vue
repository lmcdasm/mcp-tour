<template>
  <q-page padding class="q-gutter-md">
    <!-- Top Panel -->
    <q-card class="full-width">
      <q-card-section>
        <div class="text-h6">Defined McpServers</div>
        <div v-for="server in servers" :key="server.id" class="row items-center q-mb-sm">
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

    <!-- Toolbar -->
    <div class="row q-gutter-sm q-my-md">
      <div draggable="true" @dragstart="startDrag('server')">
        <q-btn flat icon="memory" label="New McpServer" />
      </div>
      <div draggable="true" @dragstart="startDrag('prompt')">
        <q-btn flat icon="description" label="Prompt" :disable="!hasServer" />
      </div>
      <div draggable="true" @dragstart="startDrag('resource')">
        <q-btn flat icon="inventory_2" label="Resource" :disable="!hasServer" />
      </div>
      <div draggable="true" @dragstart="startDrag('tool')">
        <q-btn flat icon="build_circle" label="Tool" :disable="!hasServer" />
      </div>
      <div draggable="true" @dragstart="startDrag('link')">
        <q-btn flat icon="share" label="Link" :disable="!hasServer" />
      </div>
    </div>

    <!-- SVG Canvas -->
    <div
      ref="canvasContainer"
      class="canvas-container bg-grey-2"
      @dragover.prevent
      @drop="handleDrop"
    >
      <svg ref="svgCanvas" width="100%" height="600px"></svg>
    </div>

    <!-- Modals -->
    <create-server-dialogue v-model="dialog.server" @confirm="handleServerCreate" />
    <generic-node-dialogue
      v-model="dialog.prompt"
      type="Prompt"
      :servers="servers"
      @confirm="(d) => handleNodeCreate('prompt', d)"
    />
    <generic-node-dialogue
      v-model="dialog.resource"
      type="Resource"
      :servers="servers"
      @confirm="(d) => handleNodeCreate('resource', d)"
    />
    <generic-node-dialogue
      v-model="dialog.tool"
      type="Tool"
      :servers="servers"
      @confirm="(d) => handleNodeCreate('tool', d)"
    />
    <create-linking-dialogue
      v-model="dialog.link"
      :servers="currentCanvas.getFirstServer ? [currentCanvas.getFirstServer] : []"
      :prompts="currentCanvas.getPrompts()"
      :resources="currentCanvas.getResources()"
      :tools="currentCanvas.getTools()"
      @confirm="handleLinkCreate"
    />
  </q-page>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import * as d3 from 'd3'
import { useCurrentCanvasStore } from 'src/stores/currentCanvasStore'

import CreateServerDialogue from 'src/components/CreateServerDialogue.vue'
import GenericNodeDialogue from 'src/components/GenericNodeDialogue.vue'
import CreateLinkingDialogue from 'src/components/CreateLinkingDialogue.vue'

const currentCanvas = useCurrentCanvasStore()
const hasServer = computed(() => currentCanvas.hasServer)

const svgCanvas = ref(null)
const draggedType = ref(null)
const lastDropCoords = ref([0, 0])
const servers = ref([])
const dialog = ref({
  server: false,
  prompt: false,
  resource: false,
  tool: false,
  link: false
})

function startDrag(type) {
  draggedType.value = type
}

function handleDrop(event) {
  const coords = d3.pointer(event, svgCanvas.value)
  lastDropCoords.value = coords

  if (draggedType.value === 'server') {
    if (svgCanvas.value && svgCanvas.value.children.length > 0) {
      const shouldClear = confirm('Canvas is not empty. Clear existing content?')
      if (!shouldClear) return
      svgCanvas.value.innerHTML = ''
      currentCanvas.reset()
    }
    dialog.value.server = true
  } else if (draggedType.value === 'link') {
    dialog.value.link = true 
  } else {
    dialog.value[draggedType.value] = true
  }
}

function handleServerCreate(data) {
  const [x, y] = lastDropCoords.value
  const svg = d3.select(svgCanvas.value)
  const group = svg.append('g').attr('transform', `translate(${x},${y})`)

  group.append('rect').attr('width', 140).attr('height', 60).attr('fill', '#ab47bc')
  group.append('text')
    .attr('x', 70)
    .attr('y', 30)
    .attr('text-anchor', 'middle')
    .attr('alignment-baseline', 'middle')
    .text(data.id)
    .style('fill', '#fff')

  currentCanvas.addNode({ type: 'server', id: data.id })
}

function handleNodeCreate(type, data) {
  const [x, y] = lastDropCoords.value
  const svg = d3.select(svgCanvas.value)
  const group = svg.append('g').attr('transform', `translate(${x},${y})`)

  group.append('circle').attr('r', 30).attr('fill', typeColor(type))
  group.append('text')
    .attr('text-anchor', 'middle')
    .attr('y', 5)
    .text(type.toUpperCase())
    .style('font-size', '10px')
    .style('fill', '#fff')

  currentCanvas.addNode({ type, id: data.id, linkTo: data.linkTo })
}

function handleLinkCreate(data) {
  const svg = d3.select(svgCanvas.value)
  const source = [...svg.selectAll('g').nodes()].find(g => g.textContent === data.from)
  const allTargets = [...(data.prompts || []), ...(data.resources || []), ...(data.tools || [])]

  allTargets.forEach(targetId => {
    const target = [...svg.selectAll('g').nodes()].find(g => g.textContent === targetId)
    if (source && target) {
      const bbox1 = source.getBBox()
      const bbox2 = target.getBBox()
      svg.insert('line', ':first-child')
        .attr('x1', bbox1.x + bbox1.width / 2)
        .attr('y1', bbox1.y + bbox1.height / 2)
        .attr('x2', bbox2.x + bbox2.width / 2)
        .attr('y2', bbox2.y + bbox2.height / 2)
        .attr('stroke', 'black')
        .attr('stroke-width', 2)

      currentCanvas.addLink({ from: data.from, to: targetId })
    }
  })
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
  d3.select(svgCanvas.value).append('rect')
    .attr('width', '100%')
    .attr('height', '100%')
    .attr('fill', 'transparent')

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
    const flush = window.__sseBuffer.filter(m => m.msg_id === 'list_server_ops')
    flush.forEach(({ payload }) => updateList(payload))
    window.__sseBuffer = window.__sseBuffer.filter(m => m.msg_id !== 'list_server_ops')
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

