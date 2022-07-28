CONFIG_PATH=${HOME}/.tls/
.PHONY: init

init:
	mkdir -p ${CONFIG_PATH}
.PHONY: gencert
gencert:
	cfssl gencert \
		-initca tls/ca-csr.json | cfssljson -bare ca
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=tls/ca-config.json \
		-profile=server \
		tls/server-csr.json | cfssljson -bare server
	cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=tls/ca-config.json \
		-profile=client \
		tls/client-csr.json | cfssljson -bare client
	mv *.pem *.csr ${CONFIG_PATH}


# ==============================================================================
# Docker

docker_dev:
	@echo Starting local docker dev compose
	docker-compose -f docker-compose.yaml up --build --remove-orphans

docker_dev_down:
	@echo stoping local docker dev compose
	docker-compose -f docker-compose.yaml down --remove-orphans

local:
	@echo Starting local docker compose
	docker-compose -f docker-compose-local.yaml up -d --build --remove-orphans

local_down:
	@echo Stoping local docker compose
	docker-compose -f docker-compose-local.yaml down --remove-orphans

make_user_service_image:
	docker build -f docker/user_service.Dockerfile -t reza879/user_service:$(git rev-parse --short HEAD) .
	docker push reza879/user_service:$(git rev-parse --short HEAD)

test_coverage:
	go test --cover ./...

user_service_proto:
	protoc \
		--go_out=user_service/proto/kafka --go_opt=paths=source_relative \
		--go-grpc_out=user_service/proto/kafka --go-grpc_opt=paths=source_relative \
		--proto_path=user_service/proto/kafka user_service/proto/kafka/user_kafka.proto

	protoc \
		--go_out=user_service/proto/grpc --go_opt=paths=source_relative \
		--go-grpc_out=user_service/proto/grpc --go-grpc_opt=paths=source_relative \
		--proto_path=user_service/proto/grpc user_service/proto/grpc/users.proto

message_service_proto:
	protoc \
		--go_out=message_service/proto/kafka --go_opt=paths=source_relative \
		--go-grpc_out=message_service/proto/kafka --go-grpc_opt=paths=source_relative \
		--proto_path=message_service/proto/kafka  message_service/proto/kafka/kafka.proto

	protoc \
		--go_out=message_service/proto/grpc --go_opt=paths=source_relative \
		--go-grpc_out=message_service/proto/grpc --go-grpc_opt=paths=source_relative \
		--proto_path=message_service/proto/grpc message_service/proto/grpc/message.proto

library_service_proto:
	protoc \
		--go_out=library_service/proto/kafka --go_opt=paths=source_relative \
		--go-grpc_out=library_service/proto/kafka --go-grpc_opt=paths=source_relative \
		--proto_path=library_service/proto/kafka  library_service/proto/kafka/kafka.proto

	protoc \
		--go_out=library_service/proto/grpc --go_opt=paths=source_relative \
		--go-grpc_out=library_service/proto/grpc --go-grpc_opt=paths=source_relative \
		--proto_path=library_service/proto/grpc library_service/proto/grpc/library.proto

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


#=====================================================
# kuberneties
k8s_install:
	helm install kafka bitnami/kafka
	helm install test-microservice deploy/test-microservice/

k8s_update:
	helm upgrade test-microservice deploy/test-microservice/

k8s_uninstall:
	helm uninstall test-microservice 
	helm uninstall kafka 

#=====================================================
# swagger
swagger:
	echo "Starting swagger generating"
	#swag init -g api_service/ports/http/**/*.go
	swag init -g **/**/*.go  --parseDependency
