protocol:
	protoc protocols/*.proto -I protocols/ --go_out=. --go-grpc_out=.

server:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o auth.server main.go

.PHONY: protocol echoserver 
