<template>
  <q-dialog v-model="dialog">
    <q-card style="min-width: 600px">
      <!-- Title -->
      <q-card-section>
        <div class="text-h6">Create Link</div>
      </q-card-section>

      <!-- Source Server -->
      <q-card-section>
        <q-card flat bordered class="bg-grey-2">
          <q-card-section>
            <div class="text-subtitle2">Source Server</div>
            <div class="text-bold q-mt-sm">
              {{ displayServer }}
            </div>
          </q-card-section>
        </q-card>
      </q-card-section>

      <!-- Target Selectors -->
      <q-card-section class="row q-col-gutter-md">
        <!-- Prompts -->
        <div class="col-4">
          <q-card flat bordered>
            <q-card-section>
              <q-checkbox v-model="enablePrompt" label="Target Prompts" />
              <q-select
                v-if="enablePrompt"
                v-model="form.prompts"
                label="Select Prompt(s)"
                :options="groupedPrompts"
                option-label="label"
                option-value="value"
                emit-value
                map-options
                multiple
                use-chips
                dense
                class="q-mt-sm"
              />
            </q-card-section>
          </q-card>
        </div>

        <!-- Resources -->
        <div class="col-4">
          <q-card flat bordered>
            <q-card-section>
              <q-checkbox v-model="enableResource" label="Target Resources" />
              <q-select
                v-if="enableResource"
                v-model="form.resources"
                label="Select Resource(s)"
                :options="groupedResources"
                option-label="label"
                option-value="value"
                emit-value
                map-options
                multiple
                use-chips
                dense
                class="q-mt-sm"
              />
            </q-card-section>
          </q-card>
        </div>

        <!-- Tools -->
        <div class="col-4">
          <q-card flat bordered>
            <q-card-section>
              <q-checkbox v-model="enableTool" label="Target Tools" />
              <q-select
                v-if="enableTool"
                v-model="form.tools"
                label="Select Tool(s)"
                :options="groupedTools"
                option-label="label"
                option-value="value"
                emit-value
                map-options
                multiple
                use-chips
                dense
                class="q-mt-sm"
              />
            </q-card-section>
          </q-card>
        </div>
      </q-card-section>

      <!-- Actions -->
      <q-card-actions align="right">
        <q-btn flat label="Cancel" @click="dialog = false" />
        <q-btn color="primary" label="Link" @click="confirm" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { useCurrentCanvasStore } from 'src/stores/currentCanvasStore'

const canvasStore = useCurrentCanvasStore()

const props = defineProps({
  modelValue: Boolean,
  servers: Array
})

const emit = defineEmits(['update:modelValue', 'confirm'])

const dialog = ref(false)
const form = ref({
  server: '',
  prompts: [],
  resources: [],
  tools: []
})

const enablePrompt = ref(false)
const enableResource = ref(false)
const enableTool = ref(false)

watch(() => props.modelValue, val => {
  dialog.value = val
  if (val) {
    console.log('[LinkingDialog] dialog opened')
    console.log('[LinkingDialog] canvasStore.nodes =', canvasStore.nodes)
    console.log('[LinkingDialog] getPrompts =', canvasStore.getPrompts)
    console.log('[LinkingDialog] getResources =', canvasStore.getResources)
    console.log('[LinkingDialog] getTools =', canvasStore.getTools)
  }
})

watch(() => canvasStore.currentContext.serverName, (serverId) => {
  if (serverId) {
    console.log('[LinkingDialog] Detected context server update:', serverId)
    form.value.server = serverId
  }
}, { immediate: true })

const displayServer = computed(() => form.value.server || 'No server found')

// Grouping helper
function groupBySource(items = []) {
  const unwrap = (proxy) => JSON.parse(JSON.stringify(proxy))

  const canvasItems = items
    .map(unwrap)
    .filter(i => i && i.source === 'canvas')
    .map(i => ({
      label: i.label || i.type?.toUpperCase() || i.id || 'unnamed',
      value: i.id
    }))

  const liveItems = items
    .map(unwrap)
    .filter(i => i && i.source === 'live')
    .map(i => ({
      label: i.label || i.type?.toUpperCase() || i.id || 'unnamed',
      value: i.id
    }))

  console.log('[groupBySource] Canvas:', canvasItems)
  console.log('[groupBySource] Live:', liveItems)

  return [
    { label: '--- Canvas ---', children: canvasItems },
    { label: '--- Live ---', children: liveItems }
  ]
}

// Wrapping to trigger reactivity properly
const prompts = computed(() => [...canvasStore.getPrompts])
const resources = computed(() => [...canvasStore.getResources])
const tools = computed(() => [...canvasStore.getTools])

const groupedPrompts = computed(() => groupBySource(prompts.value))
const groupedResources = computed(() => groupBySource(resources.value))
const groupedTools = computed(() => groupBySource(tools.value))

function confirm() {
  emit('confirm', {
    server: form.value.server,
    prompts: enablePrompt.value ? form.value.prompts : [],
    resources: enableResource.value ? form.value.resources : [],
    tools: enableTool.value ? form.value.tools : []
  })
  dialog.value = false
}
</script>

