GO111MODULE := auto

go-module:
	export GO111MODULE=$(GO111MODULE);

generate:
	protoc -I=./api/proto/v1 --go_out=plugins=grpc:pkg/api/v1 part_handler.proto

build: generate go-module
	go mod download; \
	go build -o ./bin/part_handler ./cmd/part_handler

run:
	./bin/part_handler

cover: go-module
	go test -p 1 -coverprofile=cover.out `go list ./... | grep -v mock` && grep -v mock \
	cover.out > coverage.out  &&  go tool cover -func=coverage.out
	rm -f cover.out
	rm -f coverage.out