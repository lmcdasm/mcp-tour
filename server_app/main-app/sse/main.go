package sse

import (
	"encoding/json"
	//"net/http"
	"sync"
	"io"

	"github.com/gin-gonic/gin"
)

type SSEMessage struct {
	MsgID   string      `json:"msg_id"`
	Payload interface{} `json:"payload"`
}

type SSEClient struct {
	ch chan SSEMessage
}

type SSEManager struct {
	clients    map[*SSEClient]bool
	register   chan *SSEClient
	unregister chan *SSEClient
	broadcast  chan SSEMessage
	mu         sync.Mutex
}

func NewSSEManager() *SSEManager {
	return &SSEManager{
		clients:    make(map[*SSEClient]bool),
		register:   make(chan *SSEClient),
		unregister: make(chan *SSEClient),
		broadcast:  make(chan SSEMessage, 100),
	}
}

func (m *SSEManager) Run() {
	for {
		select {
		case client := <-m.register:
			m.mu.Lock()
			m.clients[client] = true
			m.mu.Unlock()
		case client := <-m.unregister:
			m.mu.Lock()
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				close(client.ch)
			}
			m.mu.Unlock()
		case msg := <-m.broadcast:
			m.mu.Lock()
			for client := range m.clients {
				select {
				case client.ch <- msg:
				default:
					delete(m.clients, client)
					close(client.ch)
				}
			}
			m.mu.Unlock()
		}
	}
}

func (m *SSEManager) Broadcast(msg SSEMessage) {
	m.broadcast <- msg
}

func (m *SSEManager) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := &SSEClient{ch: make(chan SSEMessage, 100)}
		m.register <- client

		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-client.ch; ok {
				data, _ := json.Marshal(msg)
				c.SSEvent("message", string(data))
				return true
			}
			return false
		})

		defer func() {
			m.unregister <- client
		}()
	}
}

