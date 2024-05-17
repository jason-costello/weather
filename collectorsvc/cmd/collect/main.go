package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	"github.com/IBM/sarama"
	_ "github.com/IBM/sarama"

	collectorsvc "collector"
	"collector/tracing"
)

func main() {
	brokersCSV := flag.String("brokerAddresses", "127.0.0.1:9092", "kafka brokers ip:port.  If more than one separate with comma. defaults to 127.0.0.1:9092 if not provided")
	topic := flag.String("kt", "raw-weather-reports", "kafka topic name, defaults to raw-weather-reports if not provided")
	interval := flag.String("i", "10", "the interval, in seconds, between polling the collection URLs.\nDefault value is 60 seconds.")
	flag.Parse()

	brokers := strings.Split(*brokersCSV, ",")

	tracing.InitTracer("collectorsvc")
	var intervalDuration time.Duration
	var err error
	intervalDuration, err = time.ParseDuration(*interval + "s")
	if err != nil {
		intervalDuration, _ = time.ParseDuration("10s")
	}

	urlMap := map[collectorsvc.ReportType]string{
		collectorsvc.Wind:    "https://www.spc.noaa.gov/climo/reports/today_wind.csv",
		collectorsvc.Hail:    "https://www.spc.noaa.gov/climo/reports/today_hail.csv",
		collectorsvc.Tornado: "https://www.spc.noaa.gov/climo/reports/today_torn.csv",
	}

	collector, err := collectorsvc.NewCollector(*topic, brokers, urlMap, sarama.NewRandomPartitioner, sarama.WaitForAll, true)
	if err != nil {
		log.Fatal("failed to create the collect: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Printf("Starting collect service")
	collector.Poll(ctx, intervalDuration)

}
