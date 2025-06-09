<template>
  <q-card-section class="text-center">
    <div class="text-h6">MCP Client Count</div>
    <q-knob
      v-model="progress"
      :min="0"
      :max="100"
      show-value
      size="80px"
      color="deep-green"
      track-color="grey-3"
      class="q-mt-sm"
    />
  </q-card-section>
</template>

<script setup>
import { ref } from 'vue'

const progress = ref(0)

function setValue(val) {
  progress.value = val
}

defineExpose({ setValue })
</script>
