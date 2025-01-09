package consumer

import (
	"fmt"
)

// DataProcessorInterface is a minimal interface that defines a method to get data.
type DataProcessorInterface interface {
	GetData() string
}

// DataConsumer is a consumer that uses the DataProcessorInterface.
type DataConsumer struct {
	provider DataProcessorInterface
}

// ProcessData retrieves data from the provider and processes it.
func (dc *DataConsumer) ProcessData() {
	data := dc.provider.GetData()
	fmt.Println("Processed Data:", data)
}

// NewDataConsumer creates a new DataConsumer with a DataProvider.
func NewDataConsumer(provider DataProcessorInterface) *DataConsumer {
	return &DataConsumer{provider: provider}
}
