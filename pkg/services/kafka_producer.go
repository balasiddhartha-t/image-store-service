package services

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/Shopify/sarama"
)

func AlbumKafkaProducer(operation string) {
	//Command line arguments for the Producer to consume
	var (
		brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("kafka:9092").Strings()
		topic      = kingpin.Flag("topic", "Topic name").Default("ImageStatus").String()
		maxRetry   = kingpin.Flag("maxRetry", "Retry limit").Default("5").Int()
	)

	kingpin.Parse()
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = *maxRetry
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(*brokerList, config)
	if err != nil {
		fmt.Print(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			fmt.Print(err)
		}
	}()
	//Form the message object for the Topic
	msg := &sarama.ProducerMessage{
		Topic: *topic,
		Value: sarama.StringEncoder(operation),
	}

	//send a message to the producer
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)
}
