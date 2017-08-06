package main

import "github.com/shomali11/eventbus"

func main() {
	client := eventbus.NewClient()
	defer client.Close()
}
