<template>
  <q-card-section>
    <div class="text-h6">Discovery Ops Gauge</div>
    <q-linear-progress :value="progress" color="secondary" size="20px" />
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
