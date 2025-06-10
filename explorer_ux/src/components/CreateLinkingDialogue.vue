<template>
  <q-dialog v-model="dialog">
    <q-card style="min-width: 600px">
      <!-- Row 1: Title -->
      <q-card-section>
        <div class="text-h6">Create Link</div>
      </q-card-section>

      <!-- Row 2: Source Server -->
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

      <!-- Row 3: Targets -->
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

      <!-- Row 4: Actions -->
      <q-card-actions align="right">
        <q-btn flat label="Cancel" @click="dialog = false" />
        <q-btn color="primary" label="Link" @click="confirm" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup>
import { ref, watch, watchEffect, defineProps, defineEmits, computed } from 'vue'
import { useCurrentCanvasStore } from 'src/stores/currentCanvasStore'

const canvasStore = useCurrentCanvasStore()

const props = defineProps({
  modelValue: Boolean,
  servers: Array,
  prompts: Array,
  resources: Array,
  tools: Array
})

const emit = defineEmits(['update:modelValue', 'confirm'])

const dialog = ref(false)
const form = ref({
  server: '',
  prompts: [],
  resources: [],
  tools: []
})

// sync dialog visibility
watch(() => props.modelValue, val => (dialog.value = val))
watch(dialog, val => emit('update:modelValue', val))

// enable flags
const enablePrompt = ref(false)
const enableResource = ref(false)
const enableTool = ref(false)

watch(
  () => canvasStore.currentContext.serverName,
  (serverId) => {
    if (serverId) {
      console.log('[LinkingDialog] Detected context server update:', serverId)
      form.value.server = serverId
    }
  },
  { immediate: true }
)

watchEffect(() => {
  console.log('[LinkingDialog] form.server =', form.value.server)
  console.log('[LinkingDialog] currentContext =', canvasStore.currentContext)
})


const displayServer = computed(() => form.value.server || 'No server found')

function confirm() {
  emit('confirm', {
    server: form.value.server,
    prompts: enablePrompt.value ? form.value.prompts : [],
    resources: enableResource.value ? form.value.resources : [],
    tools: enableTool.value ? form.value.tools : []
  })
  dialog.value = false
}

// Grouping utility
function groupBySource(items) {
  return [
    {
      label: '--- Canvas ---',
      children: items.filter(i => i.source === 'canvas').map(i => ({
        label: i.id,
        value: i.id
      }))
    },
    {
      label: '--- Live ---',
      children: items.filter(i => i.source === 'live').map(i => ({
        label: i.id,
        value: i.id
      }))
    }
  ]
}

const groupedPrompts = computed(() => groupBySource(props.prompts))
const groupedResources = computed(() => groupBySource(props.resources))
const groupedTools = computed(() => groupBySource(props.tools))
</script>

