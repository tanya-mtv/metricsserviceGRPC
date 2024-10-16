metricsservice using GRPC

install utils
go install github.com/bufbuild/buf/cmd/buf@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

command line for generation proto
protoc --go_out=pkg/metricservice_v1 --go_opt=paths=source_relative --go-grpc_out=pkg/metricservice_v1  --go-grpc_opt=paths=source_relative api/metricservice_v1/service.proto


Generation proto files with buf-gen-yaml

Connection to Postgres DB
postgresql://user:password@localhost:5432/dbname?sslmode=disable

docs
https://habr.com/ru/articles/774796/
https://eliofery.github.io/blog/2024-02-03-kodogeneraciya-protobuf-fajlov-ispolzuya-plagin-buf-backend.html
https://github.com/ozonmp/omp-template-api/blob/main/buf.gen.yaml

https://habr.com/ru/articles/658769/