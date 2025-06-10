<template>
  <q-dialog v-model="dialog">
    <q-card style="min-width: 600px">
      <q-card-section>
        <div class="text-h6">Create Link</div>

        <!-- Source Server -->
        <q-select
          v-model="form.server"
          label="Source Server"
          :options="groupedServers"
          option-label="label"
          option-value="value"
          emit-value
          map-options
          class="q-mt-md"
        />

        <!-- Target Prompts -->
        <q-select
          v-model="form.prompts"
          label="Target Prompts"
          :options="groupedPrompts"
          option-label="label"
          option-value="value"
          emit-value
          map-options
          multiple
          use-chips
          class="q-mt-md"
        />

        <!-- Target Resources -->
        <q-select
          v-model="form.resources"
          label="Target Resources"
          :options="groupedResources"
          option-label="label"
          option-value="value"
          emit-value
          map-options
          multiple
          use-chips
          class="q-mt-md"
        />

        <!-- Target Tools -->
        <q-select
          v-model="form.tools"
          label="Target Tools"
          :options="groupedTools"
          option-label="label"
          option-value="value"
          emit-value
          map-options
          multiple
          use-chips
          class="q-mt-md"
        />
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat label="Cancel" @click="dialog = false" />
        <q-btn color="primary" label="Link" @click="confirm" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup>
import { ref, watch, defineProps, defineEmits, computed } from 'vue'

const emit = defineEmits(['update:modelValue', 'confirm'])
const props = defineProps({
  modelValue: Boolean,
  servers: Array,
  prompts: Array,
  resources: Array,
  tools: Array
})

const dialog = ref(props.modelValue)
watch(() => props.modelValue, val => (dialog.value = val))
watch(dialog, val => emit('update:modelValue', val))

const form = ref({
  server: null,
  prompts: [],
  resources: [],
  tools: []
})

watch(
  () => props.servers,
  (newVal) => {
    if (newVal?.length === 1 && !form.value.server) {
      form.value.server = newVal[0].id
    }
  },
  { immediate: true }
)


function confirm() {
  emit('confirm', { ...form.value })
  dialog.value = false
}

// Utility to group Canvas vs Live
function groupBySource(items) {
  return [
    {
      label: '--- Undefined (Canvas) ---',
      children: items.filter(i => i.source === 'canvas').map(i => ({
        label: i.id,
        value: i.id
      }))
    },
    {
      label: '--- Defined (Live) ---',
      children: items.filter(i => i.source === 'live').map(i => ({
        label: i.id,
        value: i.id
      }))
    }
  ]
}

// Computed groupings
const groupedServers = computed(() => groupBySource(props.servers))
const groupedPrompts = computed(() => groupBySource(props.prompts))
const groupedResources = computed(() => groupBySource(props.resources))
const groupedTools = computed(() => groupBySource(props.tools))
</script>

