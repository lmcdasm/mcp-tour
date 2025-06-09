<template>
  <q-card-section>
    <div class="text-h6 q-mb-sm">Mcp Registry Quick Operations</div>

    <div
      v-for="item in items"
      :key="item.id"
      class="row items-center justify-between q-pa-sm q-mb-xs bg-grey-1"
      style="border: 1px solid #e0e0e0; border-radius: 6px;"
    >
      <div class="text-subtitle2">{{ item.name }}</div>

      <div class="row q-gutter-sm">
        <q-btn dense flat round size="sm" icon="play_arrow" @click="start(item)" />
        <q-btn dense flat round size="sm" icon="stop" @click="stop(item)" />
        <q-btn dense flat round size="sm" icon="info" @click="info(item)" />
      </div>
    </div>
  </q-card-section>
</template>

<script setup>
import { ref } from 'vue'

const items = ref([])

function updateList(newItems) {
  items.value = newItems
}

defineExpose({ updateList })

const start = item => console.log('Start', item)
const stop = item => console.log('Stop', item)
const info = item => console.log('Info', item)
</script>

