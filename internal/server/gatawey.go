package server

import (
	"context"
	msV1 "metricsserviceGRPC/pkg/api/metricsserviceGRPC/pkg/metricservice_v1"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *server) createGatewayServer(gatewayAddr string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := msV1.RegisterMetricServiceHandlerFromEndpoint(ctx, mux, "localhost:4041", opts)
	if err != nil {
		s.log.Fatal(err)
	}

	s.log.Info("HTTP server listening at 8081")
	// http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) { fmt.Fprintf(w, "ok") })
	if err := http.ListenAndServe("localhost:8081", mux); err != nil {
		s.log.Fatal(err)
	}

}
