<!-- src/components/CreateServerDialogue.vue -->
<template>
  <q-dialog v-model="dialog">
    <q-card style="min-width: 400px">
      <q-card-section>
        <div class="text-h6">Create MCP Server</div>
        <q-input v-model="form.id" label="ID" class="q-mt-sm" />
        <q-input v-model="form.addr" label="Address" class="q-mt-sm" />
        <q-input v-model="form.version" label="Version" class="q-mt-sm" />
        <q-select
          v-model="form.transport"
          label="Transport"
          :options="['http', 'stdio', 'in-memory']"
          class="q-mt-sm"
        />
        <q-select
          v-model="form.buildType"
          label="Build Type"
          :options="['binary', 'container']"
          class="q-mt-sm"
        />
      </q-card-section>
      <q-card-actions align="right">
        <q-btn flat label="Reset" @click="reset" />
        <q-btn flat label="Cancel" @click="dialog = false" />
        <q-btn color="primary" label="OK" @click="confirm" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup>
import { ref, watch, defineEmits, defineProps } from 'vue'
import { useCurrentCanvasStore } from 'src/stores/currentCanvasStore'

const canvasStore = useCurrentCanvasStore()
const emit = defineEmits(['update:modelValue', 'confirm'])
const props = defineProps({ modelValue: Boolean })

const dialog = ref(props.modelValue)
watch(() => props.modelValue, val => dialog.value = val)
watch(dialog, val => emit('update:modelValue', val))

const form = ref({
  id: '',
  addr: '',
  version: '',
  transport: 'http',
  buildType: 'binary'
})

function reset() {
  form.value = {
    id: '',
    addr: '',
    version: '',
    transport: 'http',
    buildType: 'binary'
  }
}

function confirm() {
  const newNode = {
    id: form.value.id,
    type: 'server',
    label: form.value.id,
    data: { ...form.value },
    source: 'canvas'
  }
  canvasStore.addNode(newNode)
  canvasStore.setContext(newNode.id)
  emit('confirm', newNode)
  dialog.value = false
  reset()
}
</script>

