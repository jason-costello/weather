package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/segmentio/kafka-go"

	"collector/cmd/collector"
	"collector/tracing"
)

func main() {
	kafkaAddress := flag.String("ka", "127.0.0.1:9092", "kafka ip:port, defaults to 127.0.0.1:9092 if not provided")
	kafkaTopic := flag.String("kt", "raw-weather-reports", "kafka topic name, defaults to raw-weather-reports if not provided")
	interval := flag.String("i", "10", "the interval, in seconds, between polling the collection URLs.\nDefault value is 60 seconds.")
	flag.Parse()

	// put addresses inm a map keyed with hail, wind, tornado...use that key to pick which proto to marshal with

	tracing.InitTracer("collector-service")
	var intervalDuration time.Duration
	var err error
	intervalDuration, err = time.ParseDuration(*interval + "s")
	if err != nil {
		intervalDuration, _ = time.ParseDuration("10s")
	}

	collectionURLs := []string{
		"https://www.spc.noaa.gov/climo/reports/today_wind.csv",
		"https://www.spc.noaa.gov/climo/reports/today_hail.csv",
		"https://www.spc.noaa.gov/climo/reports/today_torn.csv",
	}

	collector := collector.NewCollector(kafka.TCP(*kafkaAddress), *kafkaTopic, &kafka.LeastBytes{}, collectionURLs)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Printf("Starting collector service")
	collector.Poll(ctx, intervalDuration)

}
