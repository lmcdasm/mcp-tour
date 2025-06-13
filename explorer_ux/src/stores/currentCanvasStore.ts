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

    getPrompts: (state) => {
      const result = state.nodes.filter(n => n.type === 'prompt')
      console.log('[CanvasStore:getPrompts]', result)
      return result
    },

    getResources: (state) => {
      const result = state.nodes.filter(n => n.type === 'resource')
      console.log('[CanvasStore:getResources]', result)
      return result
    },

    getTools: (state) => {
      const result = state.nodes.filter(n => n.type === 'tool')
      console.log('[CanvasStore:getTools]', result)
      return result
    },

    getFirstServer: (state) => {
      const server = state.nodes.find(n => n.type === 'server') || null
      console.log('[CanvasStore:getFirstServer]', server)
      return server
    }
  },

  actions: {
    addNode(node) {
      const exists = this.nodes.find(n => n.id === node.id && n.type === node.type)
      if (exists) {
        console.warn('[CanvasStore] Skipping duplicate node:', node)
        return
      }
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

