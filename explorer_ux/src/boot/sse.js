import { boot } from 'quasar/wrappers'

export default boot(() => {
  const evtSource = new EventSource('http://192.168.1.131:10010/api/stream')
  const messageQueue = []

  const expectedRefs = [
    'gauge_client', 'gauge_server', 'gauge_discovery',
    'list_client', 'list_server', 'list_discovery',
    'console'
  ]

  // Initialize refs to avoid undefined lookups
  if (!window.__refs) window.__refs = {}

  for (const key of expectedRefs) {
    if (!window.__refs[key]) {
      window.__refs[key] = { value: null }  // placeholder
    }
  }

  let flushed = false

  const areAnyRefsReady = () => {
    return expectedRefs.some(key => window.__refs[key]?.value)
  }

  const flushQueue = () => {
    if (flushed) return

    if (areAnyRefsReady()) {
      messageQueue.forEach(fn => fn())
      messageQueue.length = 0
      flushed = true
      console.debug('[SSE] Flushed buffered messages')
    } else {
      setTimeout(flushQueue, 150)
    }
  }

  const handleMessage = (event) => {
    try {
      const msg = JSON.parse(event.data)
      const { msg_id, payload } = msg

      console.debug('[SSE] Incoming message:', msg_id)

      if (msg_id.startsWith('list_')) {
        console.debug('[SSE] List Ref lookup:', window.__refs)
      }

      switch (msg_id) {
        // Gauges
        case 'gauge_client_ops':
          console.debug('[SSE] Updating client gauge:', payload)
          window.__refs?.gauge_client?.value?.setValue(payload)
          break

        case 'gauge_server_ops':
          console.debug('[SSE] Updating server gauge:', payload)
          window.__refs?.gauge_server?.value?.setValue(payload)
          break

        case 'gauge_discovery_ops':
          console.debug('[SSE] Updating discovery gauge:', payload)
          window.__refs?.gauge_discovery?.value?.setValue(payload)
          break

        // Lists
        case 'list_client_ops':
          console.debug('[SSE] Updating client list:', payload)
          window.__refs?.list_client?.value?.updateList(payload)
          break

        case 'list_server_ops':
          console.debug('[SSE] Updating server list:', payload)
          window.__refs?.list_server?.value?.updateList(payload)
          break

        case 'list_discovery_ops':
          console.debug('[SSE] Updating discovery list:', payload)
          window.__refs?.list_discovery?.value?.updateList(payload)
          break

        // Logs
        case 'log_info':
          console.debug('[SSE] Info log:', payload)
          window.__refs?.console?.value?.addLog(payload, 'info')
          break

        case 'log_warn':
          console.debug('[SSE] Warning log:', payload)
          window.__refs?.console?.value?.addLog(payload, 'warn')
          break

        case 'log_error':
          console.debug('[SSE] Error log:', payload)
          window.__refs?.console?.value?.addLog(payload, 'error')
          break

        default:
          console.warn('[SSE] Unhandled msg_id:', msg_id)
      }
    } catch (err) {
      console.error('[SSE] Invalid event received:', err)
    }
  }

  evtSource.onmessage = (event) => {
    if (!areAnyRefsReady()) {
      console.warn('[SSE] Refs not ready. Buffering message.')
      messageQueue.push(() => handleMessage(event))
      flushQueue()
    } else {
      handleMessage(event)
    }
  }

  evtSource.onerror = (err) => {
    console.warn('[SSE] Connection error', err)
  }
})

