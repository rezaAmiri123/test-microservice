# ==============================================================================
# Docker

docker_dev:
	@echo Starting local docker dev compose
	docker-compose -f docker-compose.yaml up --build

local:
	@echo Starting local docker compose
	docker-compose -f docker-compose-local.yaml up -d --build

local_down:
	@echo Stoping local docker compose
	docker-compose -f docker-compose-local.yaml down

test_coverage:
	go test --cover ./...

user_service_proto:
	protoc \
		--go_out=user_service/proto/kafka --go_opt=paths=source_relative \
		--go-grpc_out=user_service/proto/kafka --go-grpc_opt=paths=source_relative \
		--proto_path=user_service/proto/kafka user_service/proto/kafka/kafka.proto

	protoc \
		--go_out=user_service/proto/grpc --go_opt=paths=source_relative \
		--go-grpc_out=user_service/proto/grpc --go-grpc_opt=paths=source_relative \
		--proto_path=user_service/proto/grpc user_service/proto/grpc/users.proto

message_service_proto:
	protoc \
		--go_out=message_service/proto/kafka --go_opt=paths=source_relative \
		--go-grpc_out=message_service/proto/kafka --go-grpc_opt=paths=source_relative \
		--proto_path=message_service/proto/kafka  message_service/proto/kafka/kafka.proto

# ==============================================================================
# Go migrate postgresql

force:
	migrate -database postgres://postgres:postgres@localhost:5432/auth_db?sslmode=disable -path migrations force 1

version:
	migrate -database postgres://postgres:postgres@localhost:5432/auth_db?sslmode=disable -path migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@localhost:5432/auth_db?sslmode=disable -path migrations up 1

migrate_down:
	migrate -database postgres://postgres:postgres@localhost:5432/auth_db?sslmode=disable -path migrations down 1


