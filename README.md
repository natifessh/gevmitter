# eventemitter

A lightweight, generic, thread-safe event emitter library for Go.

## Installation

```bash
go get github.com/nati_fessh/eventemitter
```

## Usage

```go
package main

import (
    "fmt"
    emitter "todo!!!"
)

type User struct {
    Name string
    Age  int
}

func main() {
    e := emitter.NewEmitter[User]()

    e.On("user_created", func(user User) error {
        fmt.Println("user created:", user.Name)
        return nil
    })

    e.Emit("user_created", User{Name: "natnael", Age: 25})
}
```

## API

### `NewEmitter[T any]() *EventEmitter[T]`
Creates a new event emitter for any type.

### `On(event string, fn HandlerFunc[T])`
Registers a handler for an event. Multiple handlers can be registered for the same event.

```go
e.On("login", func(user User) error {
    fmt.Println("user logged in:", user.Name)
    return nil
})
```

### `Emit(event string, data T) []error`
Fires all handlers for an event synchronously. Returns a slice of errors from handlers.

```go
errs := e.Emit("login", User{Name: "natnael"})
```

### `EmitAsync(event string, data T)`
Fires all handlers for an event concurrently in separate goroutines.

```go
e.EmitAsync("login", User{Name: "natnael"})
```

### `Once(event string, fn HandlerFunc[T])`
Registers a handler that fires only once then removes itself.

```go
e.Once("welcome", func(user User) error {
    fmt.Println("welcome!", user.Name)
    return nil
})
```

### `Off(event string)`
Removes all handlers for an event.

```go
e.Off("login")
```

## Features

- Generic — works with any type
- Thread safe with `sync.RWMutex`
- Sync and async emit
- Once handler support
- Zero dependencies
