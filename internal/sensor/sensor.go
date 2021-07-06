package sensor

type Service struct {}

// Metric - a sensor metric
type Metric struct {
	Name string
	Value uint
}

// NewService - returns a new sensor service
func NewService() *Service {
	return &Service{}
}

// GetMetrics - returns all metrics for the given ID
func (service *Service) GetMetrics(ID uint) ([]Metric, error) {
	// TODO: This should maybe take a date range as well
	var metrics []Metric
	metrics = append(metrics, Metric{
		Name: "Humidity",
		Value: 789,
	})
	return metrics, nil
}