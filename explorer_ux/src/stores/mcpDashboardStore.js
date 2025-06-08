import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMcpDashboardStore = defineStore('mcpDashboard', () => {
  const clientGauge = ref(0)
  const serverGauge = ref(0)
  const discoveryGauge = ref(0)

  const clientOpsList = ref([])
  const serverOpsList = ref([])
  const discoveryOpsList = ref([])

  const logs = ref([])

  function setClientGauge(val) { clientGauge.value = val }
  function setServerGauge(val) { serverGauge.value = val }
  function setDiscoveryGauge(val) { discoveryGauge.value = val }

  function setClientOpsList(list) { clientOpsList.value = list }
  function setServerOpsList(list) { serverOpsList.value = list }
  function setDiscoveryOpsList(list) { discoveryOpsList.value = list }

  function addLog(text, level = 'info') {
    logs.value.push({ text, level })
    localStorage.setItem('mcp-console-logs', JSON.stringify(logs.value))
  }

  return {
    clientGauge, serverGauge, discoveryGauge,
    clientOpsList, serverOpsList, discoveryOpsList,
    logs,
    setClientGauge, setServerGauge, setDiscoveryGauge,
    setClientOpsList, setServerOpsList, setDiscoveryOpsList,
    addLog
  }
})

