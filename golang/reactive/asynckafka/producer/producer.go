package main

import (
	"log"

	"github.com/Shopify/sarama"
)

// ProcessResponse grabs results and errors from a producer asynchronously
func ProcessResponse(producer sarama.AsyncProducer) {
	for {
		select {
		case result := <-producer.Successes():
			log.Printf("> message: \"%s\" sent to partition %d at offset %d\n", result.Value, result.Partition, result.Offset)
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
		}
	}
}
