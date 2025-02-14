package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create JetStream context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// Define the stream name and subject
	// streamName := "MY_STREAM"
	subject := "greetings"

	// Create a durable subscription
	subscription, err := js.SubscribeSync(subject, nats.Durable("my-durable-subscription"))
	if err != nil {
		log.Fatal(err)
	}

	// Fetch messages from the stream
	fmt.Println("Listening for messages...")
	for {
		msg, err := subscription.NextMsg(time.Hour * 100000)
		if err != nil {
			log.Fatal(err)
		}

		// Print the message received
		fmt.Println("Received message:", string(msg.Data))

		// Acknowledge the message
		msg.Ack()
	}
}
