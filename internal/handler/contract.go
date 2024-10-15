package handler

import (
	"metricsserviceGRPC/internal/models"
	"metricsserviceGRPC/internal/repository"
)

type metricStorage interface {
	UpdateCounter(n string, v int64) (repository.Counter, error)
	UpdateGauge(n string, v float64) (repository.Gauge, error)
	GetAll() []models.Metrics
	GetCounter(metricName string) (repository.Counter, bool)
	GetGauge(metricName string) (repository.Gauge, bool)
	UpdateMetrics([]*models.Metrics) ([]*models.Metrics, error)
}
