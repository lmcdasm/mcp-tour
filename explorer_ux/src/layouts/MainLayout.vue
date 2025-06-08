<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title>
          MCP Exploration App
        </q-toolbar-title>

        <div>MCP-Tour v0.0.2</div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
    >
      <q-list>
        <q-item-label
          header
        >
          Essential Links
        </q-item-label>

        <EssentialLink
          v-for="link in linksList"
          :key="link.title"
          v-bind="link"
        />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref } from 'vue'
import EssentialLink from 'components/EssentialLink.vue'

const linksList = [
  {
    title: 'Dashboard',
    caption: 'Main Activity Page',
    icon: 'dashboard',
    link: '/'
  },
  {
    title: 'MCP Client Operations',
    caption: 'MCP Client Side Implementations',
    icon: 'phone_android',
    link: '/client_ops'
  },
  {
    title: 'MCP Server Operations',
    caption: 'MCP Server Side Implementations',
    icon: 'usb',
    link: '/server_ops'
  },
  {
    title: 'MCP Registry/Discovery Operations',
    caption: 'MCP Registry and Discovery Implementations',
    icon: 'share',
    link: '/discovery_ops'
  },
  {
    title: 'Settings',
    caption: 'Configure Application Settings',
    icon: 'settings_applications',
    link: '/app_settings'
  }
]

const leftDrawerOpen = ref(false)

function toggleLeftDrawer () {
  leftDrawerOpen.value = !leftDrawerOpen.value
}
</script>
