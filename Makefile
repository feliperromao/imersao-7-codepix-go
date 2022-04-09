up:
	docker-compose up -d

bash:
	docker exec -ti codepix-app-1 bash

proto:
	docker exec -ti codepix-app-1 protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto

grpc-client:
	docker exec -ti codepix-app-1 evans -r repl