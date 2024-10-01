package metrics

import (
	"metricsserviceGRPC/internal/models"
	"metricsserviceGRPC/internal/repository"
)

type metricCollector interface {
	SetValueGauge(metricName string, value repository.Gauge)
	SetValueCounter(metricName string, value repository.Counter)
	GetAllCounter() map[string]repository.Counter
	GetAllGauge() map[string]repository.Gauge
	GetAllMetricsList() []models.Metrics
}
