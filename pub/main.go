package main

import (
	"fmt"
	"log"

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

	// Define stream name and subject
	streamName := "MY_STREAM"
	subject := "greetings"

	// Try to create a stream if it doesn't exist
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{subject},
	})
	if err != nil {
		// If the stream already exists, print the error (but don't fail)
		fmt.Println("Stream already exists or failed to create:", err)
	} else {
		fmt.Println("Stream created:", streamName)
	}

	// Message to be published
	message := "Hello from JetStream - Anurag"

	// Publish message to the stream
	_, err = js.Publish(subject, []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message published:", message)

	// Flush the connection to ensure the message is sent
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	// Close connection
	fmt.Println("Connection closed")
}
