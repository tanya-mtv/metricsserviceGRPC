package metrics

import (
	"context"
	"fmt"
	"math/rand"
	"metricsserviceGRPC/internal/config"
	"metricsserviceGRPC/internal/logger"
	"metricsserviceGRPC/internal/models"
	"metricsserviceGRPC/internal/repository"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type counter struct {
	num *int32
	sync.RWMutex
}

func (c *counter) inc() {
	atomic.AddInt32(c.num, 1)
}

func (c *counter) value() int32 {
	return atomic.LoadInt32(c.num)
}

func (c *counter) nulValue() {
	atomic.StoreInt32(c.num, 0)
}

type ServiceMetrics struct {
	cfg       *config.ConfigAgent
	collector metricCollector
	counter   *counter
	log       logger.Logger
}

func (sm *ServiceMetrics) MetricsMonitor() {

	var rtm runtime.MemStats

	sm.counter.inc()

	runtime.ReadMemStats(&rtm)
	sm.collector.SetValueGauge("Alloc", repository.Gauge(rtm.Alloc))
	sm.collector.SetValueGauge("BuckHashSys", repository.Gauge(rtm.BuckHashSys))
	sm.collector.SetValueGauge("Frees", repository.Gauge(rtm.Frees))
	sm.collector.SetValueGauge("GCCPUFraction", repository.Gauge(rtm.GCCPUFraction))
	sm.collector.SetValueGauge("GCSys", repository.Gauge(rtm.GCSys))
	sm.collector.SetValueGauge("HeapAlloc", repository.Gauge(rtm.HeapAlloc))
	sm.collector.SetValueGauge("HeapIdle", repository.Gauge(rtm.HeapIdle))
	sm.collector.SetValueGauge("HeapInuse", repository.Gauge(rtm.HeapInuse))
	sm.collector.SetValueGauge("HeapObjects", repository.Gauge(rtm.HeapObjects))
	sm.collector.SetValueGauge("HeapReleased", repository.Gauge(rtm.HeapReleased))
	sm.collector.SetValueGauge("HeapSys", repository.Gauge(rtm.HeapSys))
	sm.collector.SetValueGauge("LastGC", repository.Gauge(rtm.LastGC))
	sm.collector.SetValueGauge("Lookups", repository.Gauge(rtm.Lookups))
	sm.collector.SetValueGauge("MCacheInuse", repository.Gauge(rtm.MCacheInuse))
	sm.collector.SetValueGauge("MCacheSys", repository.Gauge(rtm.MCacheSys))
	sm.collector.SetValueGauge("MSpanInuse", repository.Gauge(rtm.MSpanInuse))
	sm.collector.SetValueGauge("MSpanSys", repository.Gauge(rtm.MSpanSys))
	sm.collector.SetValueGauge("Mallocs", repository.Gauge(rtm.Mallocs))
	sm.collector.SetValueGauge("NextGC", repository.Gauge(rtm.NextGC))
	sm.collector.SetValueGauge("NumForcedGC", repository.Gauge(rtm.NumForcedGC))
	sm.collector.SetValueGauge("NumGC", repository.Gauge(rtm.NumGC))
	sm.collector.SetValueGauge("OtherSys", repository.Gauge(rtm.OtherSys))
	sm.collector.SetValueGauge("PauseTotalNs", repository.Gauge(rtm.PauseTotalNs))
	sm.collector.SetValueGauge("StackInuse", repository.Gauge(rtm.StackInuse))
	sm.collector.SetValueGauge("StackSys", repository.Gauge(rtm.StackSys))
	sm.collector.SetValueGauge("Sys", repository.Gauge(rtm.Sys))
	sm.collector.SetValueGauge("TotalAlloc", repository.Gauge(rtm.TotalAlloc))

	sm.collector.SetValueCounter("PollCount", repository.Counter(sm.counter.value()))
	sm.collector.SetValueGauge("RandomValue", repository.Gauge(float64(rand.Float64())))

}

func (sm *ServiceMetrics) MetricsMonitorGopsutil(ctx context.Context) {

	memstat, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		sm.log.Error("Can't get memstat")
		return
	}
	cpustat, err := cpu.PercentWithContext(ctx, 0, true)
	if err != nil {
		sm.log.Error("Can't get cpustat")
		return
	}

	sm.collector.SetValueGauge("TotalMemory", repository.Gauge(float64(memstat.Total)))
	sm.collector.SetValueGauge("FreeMemory", repository.Gauge(float64(memstat.Total)))
	for i, val := range cpustat {
		sm.collector.SetValueGauge(fmt.Sprintf("CPUutilization%d", i+1), repository.Gauge(float64(val)))
	}

}

func newMetric(metricName, metricsType string) *models.Metrics {

	return &models.Metrics{
		ID:    metricName,
		MType: metricsType,
	}
}

func (sm *ServiceMetrics) GetAllMetricList() []models.Metrics {
	return sm.collector.GetAllMetricsList()
}
