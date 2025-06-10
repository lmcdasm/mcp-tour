// src/stores/currentCanvasStore.ts
import { defineStore } from 'pinia'

export const useCurrentCanvasStore = defineStore('currentCanvas', {
  state: () => ({
    nodes: [],
    links: [],
    currentContext: {
      serverName: null
    }
  }),

  getters: {
    hasServer: (state) => state.nodes.some(n => n.type === 'server'),
    getPrompts: (state) => () => state.nodes.filter(n => n.type === 'prompt'),
    getResources: (state) => () => state.nodes.filter(n => n.type === 'resource'),
    getTools: (state) => () => state.nodes.filter(n => n.type === 'tool'),
    getFirstServer: (state) => () => state.nodes.find(n => n.type === 'server') || null
  },

  actions: {
    addNode(node) {
      console.log('[CanvasStore] Adding node:', node)
      this.nodes.push(node)
    },
    addLink(link) {
      console.log('[CanvasStore] Adding link:', link)
      this.links.push(link)
    },
    setContext(serverId: string) {
      console.log('[CanvasStore] Setting currentContext.serverName:', serverId)
      this.currentContext.serverName = serverId
    },
    reset() {
      console.log('[CanvasStore] Resetting canvas')
      this.nodes = []
      this.links = []
      this.currentContext.serverName = null
    }
  }
})

