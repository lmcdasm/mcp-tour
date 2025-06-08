<template>
  <q-card-section class="bg-black text-white column no-wrap" style="font-family: monospace; height: 100%;">
    <div class="row items-center q-mb-sm">
      <q-btn dense flat label="Clear" icon="delete" color="red" @click="clearLogs" />
      <q-space />
      <q-btn-toggle
        v-model="filter"
        toggle-color="primary"
        :options="[
          { label: 'All', value: 'all' },
          { label: 'Info', value: 'info' },
          { label: 'Warn', value: 'warn' },
          { label: 'Error', value: 'error' }
        ]"
        dense unelevated glossy
      />
    </div>

    <div ref="logContainer" class="scroll" style="overflow-y: auto; flex: 1;">
      <div
        v-for="(msg, index) in filteredMessages"
        :key="index"
        :class="severityClass(msg)"
      >
        {{ msg.text }}
      </div>
    </div>
  </q-card-section>
</template>

<script setup>
import { ref, computed, watch, nextTick, onMounted } from 'vue'

const STORAGE_KEY = 'mcp-console-logs'
const filter = ref('all')
const logContainer = ref(null)

const messages = ref([])

// Load persisted logs
onMounted(() => {
  const saved = localStorage.getItem(STORAGE_KEY)
  if (saved) messages.value = JSON.parse(saved)
})

// Persist logs
watch(messages, (val) => {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(val))
})

// Add log entry (external or local)
const addLog = (text, level = 'info') => {
  messages.value.push({ text, level })
}

// Auto-scroll
watch(messages, async () => {
  await nextTick()
  if (logContainer.value) {
    logContainer.value.scrollTop = logContainer.value.scrollHeight
  }
})

const clearLogs = () => {
  messages.value = []
  localStorage.removeItem(STORAGE_KEY)
}

const filteredMessages = computed(() => {
  return filter.value === 'all'
    ? messages.value
    : messages.value.filter(msg => msg.level === filter.value)
})

const severityClass = (msg) => {
  switch (msg.level) {
    case 'info': return 'text-blue';
    case 'warn': return 'text-orange';
    case 'error': return 'text-red';
    default: return '';
  }
}

defineExpose({ addLog })
</script>

