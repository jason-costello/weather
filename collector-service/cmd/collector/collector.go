package collector

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/segmentio/kafka-go"
)

type Provider interface {
	Publish(ctx context.Context, body []byte) error
}

type Collector struct {
	kafkaWriter    *kafka.Writer
	collectionURLs []string
}

func NewCollector(addr net.Addr, topic string, balancer kafka.Balancer, collectionURLs []string) *Collector {
	return &Collector{
		kafkaWriter: &kafka.Writer{
			Addr:     addr,
			Topic:    topic,
			Balancer: balancer,
		},
		collectionURLs: collectionURLs,
	}
}

func (c *Collector) Collect(ctx context.Context) error {
	var errs error
	for _, u := range c.collectionURLs {
		log.Println("collecting from: ", u)
		resp, err := http.Get(u)
		if err != nil {
			errors.Join(errs, err)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			errors.Join(errs, fmt.Errorf("failed to collect latest report from %q. response status: %s", u, resp.Status))
			continue
		}
		defer resp.Body.Close()

		bdy, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		if err := c.Publish(ctx, bdy); err != nil {
			fmt.Println("errors: ", err)
			continue
		}
		log.Println("Item published: ", string(bdy))
	}
	return nil
}

func (c *Collector) Publish(ctx context.Context, body []byte) error {
	return c.kafkaWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte(fmt.Sprintf("Key-%d", time.Now().Unix())),
		Value: body,
		Time:  time.Time{},
	})
}

func (c *Collector) Poll(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := c.Collect(ctx); err != nil {
				log.Println("failed to collect: ", err)
			}

		}
	}
}
