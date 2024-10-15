package server

import (
	"context"
	"log"
	"metricsserviceGRPC/internal/config"
	"metricsserviceGRPC/internal/handler"
	"metricsserviceGRPC/internal/logger"
	"metricsserviceGRPC/internal/models"
	"metricsserviceGRPC/internal/repository"
	msV1 "metricsserviceGRPC/pkg/api/metricsserviceGRPC/pkg/metricservice_v1"

	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type metricStorage interface {
	UpdateCounter(n string, v int64) (repository.Counter, error)
	UpdateGauge(n string, v float64) (repository.Gauge, error)
	GetAll() []models.Metrics
	GetCounter(metricName string) (repository.Counter, bool)
	GetGauge(metricName string) (repository.Gauge, bool)
	UpdateMetrics([]*models.Metrics) ([]*models.Metrics, error)
}

type server struct {
	cfg  *config.ConfigServer
	log  logger.Logger
	stor metricStorage
}

func NewServer(cfg *config.ConfigServer, log logger.Logger) *server {
	return &server{
		cfg: cfg,
		log: log,
	}
}

func (s *server) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	db, err := repository.NewPostgresDB(s.cfg.DSN)
	if err != nil {
		s.log.Fatal("Failed to initialaze db: %s", err.Error())
	} else {
		s.log.Info("Success connection to db")
		defer db.Close()
	}

	s.stor = repository.NewDBStorage(db, s.log)
	handl := handler.NewGRPCServer(s.stor, s.cfg, s.log)

	lis, err := net.Listen("tcp", "localhost:4041")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServ := grpc.NewServer(opts...)

	msV1.RegisterMetricServiceServer(grpcServ, handl)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServ)

	go func() {
		s.log.Info("Starting Server...")
		if err := grpcServ.Serve(lis); err != nil {
			s.log.Fatal("failed to serve ", err)
		}
	}()

	<-ctx.Done()
	stop()
	return nil
}
