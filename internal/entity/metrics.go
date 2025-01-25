package entity

type Metrics struct {
	Metrics map[string]string
}

func NewMetrics(metricData map[string]string) (*Metrics, error) {
	return &Metrics{
		Metrics: metricData,
	}, nil
}
