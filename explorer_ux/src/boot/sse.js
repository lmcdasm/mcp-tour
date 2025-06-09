import { boot } from 'quasar/wrappers'
//import { useMcpDashboardStore } from 'stores/mcpDashboardStore'

export default boot(() => {
  // const store = useMcpDashboardStore()
  const evtSource = new EventSource('http://192.168.1.131:10010/api/stream')

  evtSource.onmessage = (event) => {
    try {
      const msg = JSON.parse(event.data)
      const { msg_id, payload } = msg

      switch (msg_id) {
        // Gauges
        case 'gauge_client_ops':
          window.__refs?.gauge_client?.value?.setValue(payload)
          break
        case 'gauge_server_ops':
          window.__refs?.gauge_server?.value?.setValue(payload)
          break
        case 'gauge_discovery_ops':
          window.__refs?.gauge_discovery?.value?.setValue(payload)
          break

        // Lists
        case 'list_client_ops':
          window.__refs?.list_client?.value?.updateList(payload)
          break
        case 'list_server_ops':
          window.__refs?.list_server?.value?.updateList(payload)
          break
        case 'list_discovery_ops':
          window.__refs?.list_discovery?.value?.updateList(payload)
          break

        // Logs
        case 'log_info':
          window?.__mcpConsole?.value?.addLog(payload, 'info')
          break
        case 'log_warn':
          window?.__mcpConsole?.value?.addLog(payload, 'warn')
          break
        case 'log_error':
          window?.__mcpConsole?.value?.addLog(payload, 'error')
          break
      }
    } catch (err) {
      console.error('[SSE] Invalid event received:', err)
    }
  }

  evtSource.onerror = (err) => {
    console.warn('[SSE] Connection error', err)
  }
})

