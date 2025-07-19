 
GOPATH=$(shell go env GOPATH)
GOROOT=$(shell go env GOROOT)

hello: 
	echo 'HelloWorld! GOPATH: $(GOPATH)  GOROOT: $(GOROOT)'

format:
	gofmt -w .

format_test:
	gofmt -l .

go_clear:
	go mod tidy

go_check:
	go mod verify

run:
	go run cmd/main.go

migrate:
	go run internal/migrations/migrations.go

build:
	go build -o ./bin/mail-phone-auth ./cmd/main.go

build_windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/mail-phone-auth.exe ./cmd/main.go

docker_run: 
	docker compose up -d

docker_show:
	docker ps

docker_info:
	docker info

docker_stop:
	docker compose down

swagger_version:
	$(GOPATH)/bin/swag -v

swagger_docs:
	$(GOPATH)/bin/swag fmt
	$(GOPATH)/bin/swag init -g ./internal/api/api.go

openapi_gen:
	$(GOPATH)/bin/oapi-codegen --config=./internal/api/openapi/config.json ./internal/api/openapi/openapi.json
