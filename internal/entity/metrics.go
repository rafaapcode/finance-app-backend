package entity

type Metrics struct {
	Metrics map[string]float64
}

func NewMetrics(metricData map[string]float64) *Metrics {
	return &Metrics{
		Metrics: metricData,
	}
}
