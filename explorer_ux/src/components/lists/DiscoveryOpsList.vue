<template>
  <q-card-section>
    <div class="text-h6">Discovery Ops</div>
    <q-list dense>
      <q-item v-for="item in items" :key="item.id">
        <q-item-section>{{ item.name }}</q-item-section>
        <q-item-section side>
          <q-btn dense flat icon="play_arrow" @click="start(item)" />
          <q-btn dense flat icon="stop" @click="stop(item)" />
          <q-btn dense flat icon="info" @click="info(item)" />
        </q-item-section>
      </q-item>
    </q-list>
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

