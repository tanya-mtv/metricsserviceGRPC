package server

import (
	"context"
	"log"
	"metricsserviceGRPC/internal/config"
	"metricsserviceGRPC/internal/handler"
	"metricsserviceGRPC/internal/logger"
	"metricsserviceGRPC/internal/repository"
	msV1 "metricsserviceGRPC/pkg/api/metricsserviceGRPC/pkg/metricservice_v1"
	"time"

	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	_ "github.com/lib/pq"
	// gw "github.com/yourorg/yourrepo/proto/gen/go/your/service/v1/your_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// type metricStorage interface {
// 	UpdateCounter(n string, v int64) (repository.Counter, error)
// 	UpdateGauge(n string, v float64) (repository.Gauge, error)
// 	GetAll() []models.Metrics
// 	GetCounter(metricName string) (repository.Counter, bool)
// 	GetGauge(metricName string) (repository.Gauge, bool)
// 	UpdateMetrics([]*models.Metrics) ([]*models.Metrics, error)
// }

type server struct {
	cfg *config.ConfigServer
	log logger.Logger
	// stor metricStorage
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

	stor := repository.NewDBStorage(db, s.log)
	handl := handler.NewGRPCServer(stor, s.cfg, s.log)

	lis, err := net.Listen("tcp", "localhost:4041")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		// grpc_prometheus.UnaryServerInterceptor,
		// grpc_opentracing.UnaryServerInterceptor(),
		grpcrecovery.UnaryServerInterceptor(),
		validateInterceptor,
	)))
	grpcServ := grpc.NewServer(opts...)

	msV1.RegisterMetricServiceServer(grpcServ, handl)
	// Register reflection service on gRPC server.
	reflection.Register(grpcServ)

	go func() {
		s.log.Info("Starting GRPC Server...")
		if err := grpcServ.Serve(lis); err != nil {
			s.log.Fatal("failed to serve ", err)
		}
	}()

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	go s.createGatewayServer("")

	time.Sleep(5 * time.Second)
	<-ctx.Done()
	stop()
	return nil
}
