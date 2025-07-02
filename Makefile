hello: 
	echo 'HelloWorld!'

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
