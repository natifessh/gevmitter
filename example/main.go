package main

import (
	"fmt"
)

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
	emitter.On("log_in", func(data User) error {
		println(data.name)
		return nil
	})
	emitter.Emit("log_in", User{name: "Meron", age: 18})
}
