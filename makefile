regenerate:
	protoc --proto_path=protopb protopb/*.proto --go_out=protopb/  --go-grpc_out=protopb/ protopb/*.proto 

run-server:
		go run cmd/client-service/main.go 

run-client:
		go run cmd/server-service/main.go

