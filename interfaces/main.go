package main

import (
	"github.com/mostafababaii/go-design-patterns/interfaces/consumer"
	"github.com/mostafababaii/go-design-patterns/interfaces/producer"
)

func main() {
	// Create an instance of DataProvider
	dataProvider := &producer.DataProvider{}

	// Create an instance of DataConsumer with the DataProvider
	dataConsumer := consumer.NewDataConsumer(dataProvider)

	// Process the data
	dataConsumer.ProcessData()
}
