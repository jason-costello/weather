package collectorsvc

//go:generate stringer -type=ReportType

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

type ReportType int

const (
	Hail ReportType = iota
	Wind
	Tornado
)

type Provider interface {
	Publish(ctx context.Context, body []byte) error
}

type Collector struct {
	topic          string
	collectionURLs map[ReportType]string
	producer       sarama.SyncProducer
}

func NewCollector(topic string, brokersAddresses []string, collectionURLs map[ReportType]string, partitioner sarama.PartitionerConstructor, acks sarama.RequiredAcks, returnPublishSuccess bool) (*Collector, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = partitioner
	config.Producer.RequiredAcks = acks
	config.Producer.Return.Successes = returnPublishSuccess
	producer, err := sarama.NewSyncProducer(brokersAddresses, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create sync producer with brokers %s: %w", strings.Join(brokersAddresses, ","), err)
	}

	return &Collector{
		topic:          topic,
		collectionURLs: collectionURLs,
		producer:       producer,
	}, nil
}

func (c *Collector) CollectAndPublish(ctx context.Context) error {
	var errs error
	for rt, rURL := range c.collectionURLs {
		log.Printf("collecting %s report from: %s\n", rt.String(), rURL)
		resp, err := http.Get(rURL)
		if err != nil {
			errors.Join(errs, err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			errors.Join(errs, fmt.Errorf("failed to collect latest %s report from %q. response status: %s", rt.String(), rURL, resp.Status))
			continue
		}
		defer resp.Body.Close()

		bdy, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		msg := prepareMessage(rt, c.topic, bdy)
		partition, offset, err := c.publish(msg)
		log.Printf("partition: %v    offset: %v\n", partition, offset)
		if err != nil {
			fmt.Println("collected report but unable to publish: ", err)
			continue
		}
		log.Println("Item published: ", string(bdy))
	}
	return nil
}

func prepareMessage(reportType ReportType, topic string, message []byte) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Headers: []sarama.RecordHeader{{
			Key:   []byte("reportType"),
			Value: []byte(reportType.String()),
		}},
		Topic:     topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(message),
		Timestamp: time.Now().UTC(),
	}
	return msg
}

func (c *Collector) publish(msg *sarama.ProducerMessage) (int32, int64, error) {
	partition, offset, err := c.producer.SendMessage(msg)
	if err != nil {
		fmt.Errorf("failed to send message: %w ", err)
	}
	return partition, offset, err
}

func (c *Collector) Poll(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := c.CollectAndPublish(ctx); err != nil {
				log.Println("failed to collect: ", err)
			}

		}
	}
}
