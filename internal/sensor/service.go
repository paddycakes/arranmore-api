package sensor

import (
	"github.com/paddycakes/arranmore-api/internal/sensor/danalto"
)

// TODO: Should I enforce creation of a single instance? No. because apiKey needs to be different per client.
// TODO: Data transformations should happen in this service
// TODO: Should maybe rename internal -> services
// TODO: Maybe move transport out to top level package..

// TODO: Think I need to pass apiKey on every invocation. Not known when the service is created.


type Service struct {
	danaltoClient *danalto.Client
}

// Metric - a sensor metric
type Metric struct {
	Name string
	Value string // uint
}

// NewService - returns a new sensor service
func NewService() *Service {
	return &Service{
		danaltoClient: danalto.NewClient(),
	}
}

// GetMetrics - returns all metrics for the given ID
func (service *Service) GetMetrics(ID uint) ([]Metric, error) {
	// TODO: This should maybe take a date range as well
	//var metrics []Metric
	//metrics = append(metrics, Metric{
	//	Name: "Temperature",
	//	Value: 28,
	//})
	//fmt.Println("Returning metrics")
	//return metrics, nil

	sensorDataList, err := service.danaltoClient.GetDeviceData("U2VhbXVzQm9ubmVyOmFycmFubW9yZUlvVA==", "a81758fffe0346cd")
	if err != nil {
		return nil, err
	}

	first := sensorDataList[0];

	// TODO: Convert to model required for FE
	var metrics []Metric
	metrics = append(metrics, Metric{
		Name: "Temperature",
		Value: first.Payload.Temperature,
	})

	// return sensorDataList, nil
	return metrics, nil
}