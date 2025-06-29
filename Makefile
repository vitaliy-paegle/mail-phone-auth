hello: 
	echo 'HelloWorld!'

run:
	go run cmd/main.go

build:
	go build -o ./bin/mail-phone-auth ./cmd/main.go

build_windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/mail-phone-auth.exe ./cmd/main.go