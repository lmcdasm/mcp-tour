<!-- src/components/GenericNodeDialogue.vue -->
<template>
  <q-dialog v-model="dialog">
    <q-card style="min-width: 400px">
      <q-card-section>
        <div class="text-h6">Create {{ type }}</div>
        <q-input v-model="form.id" label="ID" class="q-mt-sm" />
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

const emit = defineEmits(['update:modelValue', 'confirm'])
const props = defineProps({
  modelValue: Boolean,
  type: String
})

const dialog = ref(props.modelValue)
watch(() => props.modelValue, val => dialog.value = val)
watch(dialog, val => emit('update:modelValue', val))

const form = ref({ id: '' })

function reset() {
  form.value = { id: '' }
}

function confirm() {
  emit('confirm', { ...form.value })
  dialog.value = false
  reset()
}
</script>

