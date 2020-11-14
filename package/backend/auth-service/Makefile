clean:
	@go clean ./...
	@go mod tidy

build:
	cd cmd && go build -trimpath -gcflags=-trimpath=%CD% -asmflags=-trimpath=%CD% -ldflags "-s -w" main.go

test:
	@go test -race -v -p 2 -coverpkg=./... -covermode=atomic -coverprofile cover.out.tmp ./...

lint:
	@golangci-lint run ./...
	@cd pkg && golangci-lint run ./...

docker-build:
	docker build -t grpc/auth-service -f ./docker/Dockerfile .
	docker tag grpc/auth-service grpc/auth-service:1.0.0

