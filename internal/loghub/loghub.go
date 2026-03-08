package loghub

import "sync"

type Hub struct {
	subscribers []chan string
	mu          sync.Mutex
}

func New() *Hub {
	return &Hub{}
}

func (h *Hub) Publish(msg string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, sub := range h.subscribers {
		sub <- msg
	}
}

func (h *Hub) Subscribe() chan string {
	ch := make(chan string)

	h.mu.Lock()
	h.subscribers = append(h.subscribers, ch)
	h.mu.Unlock()

	return ch
}