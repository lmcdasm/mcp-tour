<template>
  <q-page class="q-pa-md">
    <q-page-container>
      <!-- Top: Gauge cards -->
      <div class="row q-col-gutter-md">
        <div class="col-12 col-md-4">
          <q-card><ClientGauge ref="clientGaugeRef" /></q-card>
        </div>
        <div class="col-12 col-md-4">
          <q-card><ServerGauge ref="serverGaugeRef" /></q-card>
        </div>
        <div class="col-12 col-md-4">
          <q-card><DiscoveryGauge ref="discoveryGaugeRef" /></q-card>
        </div>
      </div>

      <!-- Middle: List cards -->
      <div class="row q-mt-md q-col-gutter-md">
        <div class="col-12 col-md-4">
          <q-card><ClientOpsList ref="clientListRef" /></q-card>
        </div>
        <div class="col-12 col-md-4">
          <q-card><ServerOpsList ref="serverListRef" /></q-card>
        </div>
        <div class="col-12 col-md-4">
          <q-card><DiscoveryOpsList ref="discoveryListRef" /></q-card>
        </div>
      </div>

      <!-- Bottom: Console log -->
      <div class="row q-mt-md">
        <div class="col-12">
          <q-card class="bg-grey-2" style="min-height: 300px;">
            <ConsoleLog ref="consoleRef" />
          </q-card>
        </div>
      </div>
    </q-page-container>
  </q-page>
</template>

<script setup>
import { ref, onMounted } from 'vue'

import ClientGauge from 'components/gauges/ClientGauge.vue'
import ServerGauge from 'components/gauges/ServerGauge.vue'
import DiscoveryGauge from 'components/gauges/DiscoveryGauge.vue'

import ClientOpsList from 'components/lists/ClientOpsList.vue'
import ServerOpsList from 'components/lists/ServerOpsList.vue'
import DiscoveryOpsList from 'components/lists/DiscoveryOpsList.vue'

import ConsoleLog from 'components/ConsoleLog.vue'

// Refs for SSE-targeted components
const consoleRef = ref(null)
const clientGaugeRef = ref(null)
const serverGaugeRef = ref(null)
const discoveryGaugeRef = ref(null)

const clientListRef = ref(null)
const serverListRef = ref(null)
const discoveryListRef = ref(null)

onMounted(() => {
  if (!window.__refs) {
    window.__refs = {}
  }
  // Expose all references globally for boot/sse.js
  window.__refs = {
    console: consoleRef,
    gauge_client: clientGaugeRef,
    gauge_server: serverGaugeRef,
    gauge_discovery: discoveryGaugeRef,
    list_client: clientListRef,
    list_server: serverListRef,
    list_discovery: discoveryListRef
  }
})
</script>

