.phony: run

# PORT=8080


run: 
	go run ./cmd/service-pdf-compose-server/main.go


gen:
	protoc \
  	--go_out=. \
  	--go_opt=paths=source_relative \
  	--go-grpc_out=. \
 	--go-grpc_opt=paths=source_relative \
  	api/api.proto


