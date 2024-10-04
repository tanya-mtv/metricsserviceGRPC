package server

import (
	"context"
	"log"
	"metricsserviceGRPC/internal/config"
	"metricsserviceGRPC/internal/logger"
	"metricsserviceGRPC/internal/models"
	"metricsserviceGRPC/internal/repository"
	"metricsserviceGRPC/pk/metricservice_v1/api/metricservice_v1"

	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type metricStorage interface {
	UpdateCounter(n string, v int64) repository.Counter
	UpdateGauge(n string, v float64) repository.Gauge
	GetAll() []models.Metrics
	GetCounter(metricName string) (repository.Counter, bool)
	GetGauge(metricName string) (repository.Gauge, bool)
	UpdateMetrics([]*models.Metrics) ([]*models.Metrics, error)
}

type server struct {
	cfg *config.ConfigServer

	log  logger.Logger
	stor metricStorage
}

func NewServer(cfg *config.ConfigServer, log logger.Logger) *server {
	return &server{
		cfg: cfg,
		log: log,
		metricservice_v1.MetricServiceServer,
	}
}

func (s *server) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// if s.cfg.DSN != "" {
	// 	db, err := repository.NewPostgresDB(s.cfg.DSN)

	// 	if err != nil {
	// 		s.log.Info("Failed to initialaze db: %s", err.Error())
	// 	} else {
	// 		s.log.Info("Success connection to db")
	// 		defer db.Close()
	// 	}
	// 	s.openStorage(ctx, db)
	// 	s.router = s.NewRouter(db)
	// } else {
	// 	s.openStorage(ctx, nil)
	// 	s.router = s.NewRouter(nil)
	// }

	lis, err := net.Listen("tcp", "localhost:4041")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	logs.RegisterLogServiceServer(s, &GRPCServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		s.log.Info("Connect listening on port: %s", s.cfg.Port)
		if err := s.router.Run(s.cfg.Port); err != nil {

			s.log.Fatal("Can't ListenAndServe on port", s.cfg.Port)
		}
	}()

	<-ctx.Done()
	stop()
	return nil
}
