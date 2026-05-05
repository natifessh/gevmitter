package main

import (
	"sync"
)

type HandlerFunc func(data any)
type EventEmitter struct {
	events map[string][]HandlerFunc
	mu     sync.Mutex
}

func NewEmitter() EventEmitter {
	return EventEmitter{
		events: map[string][]HandlerFunc{},
	}
}
func (ev *EventEmitter) On(eventname string, fn HandlerFunc) {
	ev.mu.Lock()
	defer ev.mu.Unlock()
	ev.events[eventname] = append(ev.events[eventname], fn)
}
func (ev *EventEmitter) Emit(eventname string, data any) {
	ev.mu.Lock()
	defer ev.mu.Unlock()
	handlers := ev.events[eventname]
	for _, f := range handlers {
		f(data)
	}
}
func main() {
	println("Gevmitter")
}
