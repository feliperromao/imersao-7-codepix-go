build:
	docker-compose up -d --build

up:
	docker-compose up -d

down:
	docker-compose down

bash:
	docker exec -ti app bash

proto:
	docker exec -ti app protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto

grpc-client:
	docker exec -ti app evans -r repl

cobra-init:
	docker exec -ti app cobra-cli init
