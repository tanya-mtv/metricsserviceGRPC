metricsservice using GRPC

install utils
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

command line for generation proto
protoc --go_out=pkg/metricservice_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/metricservice_v1  --go-grpc_opt=paths=source_relative api/metricservice_v1/service.proto

docs
https://habr.com/ru/articles/774796/
!!!!!!