package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Kafka broker address
	brokerAddress := "localhost:9092"

	// Topic to read messages from
	topic := "raw-weather-reports"

	// Create a new context
	ctx := context.Background()

	// Create a new Kafka reader with the specified broker address and topic
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddress},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	defer r.Close()

	// Set the offset to the first message in the topic
	r.SetOffset(0)

	fmt.Println("Starting to read messages from the topic:", topic)

	for {
		// Read a message from the Kafka topic
		fmt.Println("reading message")
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Printf("could not read message %v", err)
			continue
		}
		// Print the message key and value
		fmt.Printf("received: %s: %s\n", string(msg.Key), string(msg.Value))

		// Optional: Sleep for a short duration to simulate processing time
		time.Sleep(10 * time.Second)
	}
}
