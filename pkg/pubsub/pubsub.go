package pubsub

import (
	"sync"
)

type Message struct {
	Queue string
	Data  string
}

type Broker struct {
	subscribers map[string][]chan string
	mutex       sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]chan string),
		mutex:       sync.RWMutex{},
	}
}

func (b *Broker) Publish(msg Message) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	if subscribers, ok := b.subscribers[msg.Queue]; ok {
		for _, subscriber := range subscribers {
			subscriber <- msg.Data
		}
	}
}

func (b *Broker) Subscribe(queue string) <-chan string {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	ch := make(chan string, 1)
	b.subscribers[queue] = append(b.subscribers[queue], ch)
	return ch
}

func (b *Broker) Close() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, subscribers := range b.subscribers {
		for _, ch := range subscribers {
			close(ch)
		}
	}

	b.subscribers = make(map[string][]chan string)
}
