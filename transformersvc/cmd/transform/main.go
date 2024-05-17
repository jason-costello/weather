package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/segmentio/kafka-go/protocol"
)

func main() {
	brokers := []string{"localhost:9092"}
	topic := "raw-weather-reports"
	ctx := context.Background()

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Fatal("Could not create consumer: ", err)
	}

	subscribe(topic, consumer)
	fmt.Println("Starting to read messages from the topic:", topic)

	for {
		// Read a message from the Kafka topic
		fmt.Println("reading message")
		msg, err := consumer.ReadMessage(ctx)
		if err != nil {
			log.Printf("could not read message %v", err)
			continue
		}
		// Print the message key and value

		// Optional: Sleep for a short duration to simulate processing time
		time.Sleep(10 * time.Second)
	}
}
func subscribe(topic string, consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions(topic) // get all partitions on the given topic
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}
	initialOffset := sarama.OffsetOldest // get offset for the oldest message on the topic

	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition(topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				messageReceived(message)
			}
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) {
	log.Println("Report Type: ", string(message.Key))

	for _, line := range strings.Split(string(message.Value), "\n") {
		log.Println("processing line: ", line)
		rptMsg := *ReportMsg{}
		protocol.Unmarshal([]byte(line))
	}
	// process message, get report type
	// marshal to protobuf
	// any logic needed
	// send to topic
}
