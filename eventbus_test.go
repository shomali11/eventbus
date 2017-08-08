package eventbus

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEventbus_Fail(t *testing.T) {
	client := NewClient()
	defer client.Close()

	client.Publish("unknown", "Unknown")

	time.Sleep(time.Second)

	client.Subscribe("unknown", func(value interface{}) {
		assert.Fail(t, "Not supposed to be reached")
	})

	time.Sleep(time.Second)
}

func TestEventbus_Success(t *testing.T) {
	client := NewClient()
	defer client.Close()

	client.Subscribe("name", func(value interface{}) {
		assert.Equal(t, value, "Raed Shomali")
	})

	time.Sleep(time.Second)

	client.Publish("name", "Raed Shomali")

	time.Sleep(time.Second)
}

func TestEventbus_Multi(t *testing.T) {
	client := NewClient()
	defer client.Close()

	client.Subscribe("name", func(value interface{}) {
		assert.Equal(t, value, "Raed Shomali")
	})

	client.Subscribe("name", func(value interface{}) {
		assert.Equal(t, value, "Raed Shomali")
	})

	time.Sleep(time.Second)

	client.Publish("name", "Raed Shomali")

	time.Sleep(time.Second)
}