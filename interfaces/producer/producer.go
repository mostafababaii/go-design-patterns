package producer

// DataProvider is a concrete implementation of DataProcessorInterface.
type DataProvider struct{}

// GetData retrieves and returns a string representing the data provided by the DataProvider.
func (d *DataProvider) GetData() string {
	return "Data from DataProvider"
}
