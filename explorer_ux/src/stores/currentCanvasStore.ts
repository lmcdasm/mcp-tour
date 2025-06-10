// src/stores/currentCanvasStore.ts
import { defineStore } from 'pinia'

export const useCurrentCanvasStore = defineStore('currentCanvas', {
  state: () => ({
    nodes: [],
    links: []
  }),
  getters: {
    hasServer: (state) => state.nodes.some(n => n.type === 'server'),
    getPrompts: (state) => () => state.nodes.filter(n => n.type === 'prompt'),
    getResources: (state) => () => state.nodes.filter(n => n.type === 'resource'),
    getTools: (state) => () => state.nodes.filter(n => n.type === 'tool'),
    getFirstServer: (state) => () =>
      state.nodes.find(n => n.type === 'server') || null
  },
  actions: {
    addNode(node) {
      this.nodes.push(node)
    },
    addLink(link) {
      this.links.push(link)
    },
    reset() {
      this.nodes = []
      this.links = []
    }
  }
})

