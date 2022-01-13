# ==============================================================================
# Docker

docker_dev:
	@echo Starting local docker dev compose
	docker-compose -f docker-compose.yaml up --build

local:
	@echo Starting local docker compose
	docker-compose -f docker-compose-local.yaml up -d --build

test_coverage:
	go test --cover ./...

user_service_proto:
	protoc \
		--go_out=user_service/proto/kafka --go_opt=paths=source_relative \
		--go-grpc_out=user_service/proto/kafka --go-grpc_opt=paths=source_relative \
		--proto_path=user_service/proto/kafka user_service/proto/kafka/kafka.proto
