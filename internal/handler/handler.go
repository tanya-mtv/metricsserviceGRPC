package handler

import (
	"context"
	"metricsserviceGRPC/internal/config"
	"metricsserviceGRPC/internal/logger"

	msV1 "metricsserviceGRPC/pkg/api/metricsserviceGRPC/pkg/metricservice_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	msV1.MetricServiceServer
	storage metricStorage
	cfg     *config.ConfigServer
	log     logger.Logger
}

func NewGRPCServer(storage metricStorage, cfg *config.ConfigServer, log logger.Logger) *GRPCServer {
	return &GRPCServer{
		storage: storage,
		cfg:     cfg,
		log:     log,
	}
}

func (g *GRPCServer) PostV1(ctx context.Context, req *msV1.MetricRequest) (*msV1.MetricResponce, error) {
	err := req.ValidateAll()
	if err != nil {
		return &msV1.MetricResponce{}, err
	}
	data := req.Value
	delta := data.GetDelta()
	value := float64(data.GetValue())
	metric := Metrics{
		ID:    data.GetId(),
		MType: data.GetMType(),
		Delta: &delta,
		Value: &value,
	}

	// Put data to DB
	if metric.MType == "gauge" {
		_, err := g.storage.UpdateGauge(metric.ID, *metric.Value)
		if err != nil {
			g.log.Info("Cannot put gauge data to DB", err)
			return nil, status.Error(codes.Canceled, err.Error())
		}
	} else {
		_, err := g.storage.UpdateCounter("counter", *metric.Delta)
		if err != nil {
			g.log.Info("Cannot put counter data to DB", err)
			return nil, status.Error(codes.Canceled, err.Error())
		}

	}

	res := &msV1.MetricResponce{
		Status: codes.OK.String(),
	}
	return res, nil

}
