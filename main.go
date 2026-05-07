package main

import (
	"fmt"
	"sync"
)

type HandlerFunc[T any] func(data T) error
type EventEmitter[T any] struct {
	events map[string][]HandlerFunc[T]
	mu     sync.RWMutex
}

func NewEmitter[T any]() *EventEmitter[T] {
	return &EventEmitter[T]{
		events: make(map[string][]HandlerFunc[T]),
	}
}
func (ev *EventEmitter[T]) On(eventname string, fn HandlerFunc[T]) {
	ev.mu.Lock()
	defer ev.mu.Unlock()
	ev.events[eventname] = append(ev.events[eventname], fn)
}
func (ev *EventEmitter[T]) Emit(eventname string, data T) []error {
	ev.mu.RLock()
	handlers := make([]HandlerFunc[T], len(ev.events[eventname]))
	copy(handlers, ev.events[eventname])
	ev.mu.RUnlock()
	var errs []error
	for _, f := range handlers {
		if err := f(data); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
func (ev *EventEmitter[T]) Off(eventname string) {
	ev.mu.Lock()
	defer ev.mu.Unlock()
	delete(ev.events, eventname)
}
func (ev *EventEmitter[T]) Once(eventname string, fn HandlerFunc[T]) {
	var once sync.Once
	ev.On(eventname, func(data T) error {
		once.Do(func() {
			fn(data)
			ev.Off(eventname)
		})

		return nil
	})

}
func (ev *EventEmitter[T]) EmitAsync(eventname string, data T) {
	ev.mu.RLock()
	handlers := make([]HandlerFunc[T], len(ev.events[eventname]))
	copy(handlers, ev.events[eventname])
	ev.mu.RUnlock()
	for _, fn := range handlers {
		go fn(data)
	}
}

type User struct {
	name string
	age  int
}

func main() {
	emitter := NewEmitter[User]()
	emitter.On("user_created", func(user User) error {
		fmt.Println("hello user:", user.name, user.age)
		return nil
	})

	emitter.Once("say_hi", func(data User) error {
		fmt.Println("say hi man")
		return nil
	})
	emitter.Emit("say_hi", User{name: "man", age: 12})
	emitter.Emit("user_created", User{name: "johnny", age: 12})
}
