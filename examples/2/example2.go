package main

import (
	"fmt"
	"github.com/shomali11/eventbus"
	"time"
)

func main() {
	client := eventbus.NewClient()
	defer client.Close()

	client.Publish("test", "test")

	client.Subscribe("name", func(value interface{}) {
		fmt.Println(value)
	})

	client.Subscribe("name", func(value interface{}) {
		fmt.Println(value, "is Awesome")
	})

	client.Publish("name", "Raed Shomali")

	time.Sleep(time.Second)
}
