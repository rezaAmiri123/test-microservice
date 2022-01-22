we have a user service that handles the user work like register and login


### Jaeger UI:

http://localhost:16686

### Prometheus UI:

http://localhost:9090

### Grafana UI:

http://localhost:3000

### Swagger UI:

http://localhost:5001/swagger/index.html

```text
.
├── docker
│   └── user_service.Dockerfile
├── docker-compose-local.yaml
├── docker-compose.yaml
├── go.mod
├── go.sum
├── library_service
│   ├── adapters
│   │   ├── article_gorm_repository.go
│   │   └── pg
│   │       ├── article_repository.go
│   │       └── sql_queries.go
│   ├── agent
│   │   ├── agent.go
│   │   ├── application.go
│   │   ├── auth_client.go
│   │   ├── http_server.go
│   │   ├── logger.go
│   │   ├── mertic.go
│   │   └── tracing.go
│   ├── app
│   │   ├── app.go
│   │   ├── commands
│   │   │   └── create_article.go
│   │   └── queries
│   │       └── get_article_by_slug.go
│   ├── cmd
│   │   └── library
│   │       └── main.go
│   ├── domain
│   │   └── article
│   │       ├── article.go
│   │       ├── mock
│   │       │   └── repository.go
│   │       └── repository.go
│   ├── metrics
│   │   └── metrics.go
│   └── ports
│       └── http
│           ├── http_handlers.go
│           └── http_server.go
├── Makefile
├── migrations
│   ├── 01_create_initial_tables.down.sql
│   └── 01_create_initial_tables.up.sql
├── monitoring
│   └── prometheus_docker.yml
├── pkg
│   ├── auth
│   │   ├── auth.go
│   │   └── middleware.go
│   ├── converter
│   │   └── converter.go
│   ├── db
│   │   └── postgres
│   │       └── db_conn.go
│   ├── kafka
│   │   ├── client.go
│   │   ├── config.go
│   │   ├── constants.go
│   │   ├── consumer_group.go
│   │   ├── producer.go
│   │   ├── reader.go
│   │   └── writer.go
│   ├── logger
│   │   ├── applogger
│   │   │   └── applogger.go
│   │   └── logger.go
│   ├── redis
│   │   └── redis.go
│   ├── tracing
│   │   └── jaeger.go
│   └── utils
│       └── validator.go
├── README.md
├── test
│   ├── kafka1
│   │   ├── consumer.go
│   │   └── producer.go
│   └── main.go
└── user_service
    ├── adapters
    │   ├── user_gorm_repository.go
    │   ├── user_gorm_repository_test.go
    │   └── user_memory_repository.go
    ├── agent
    │   ├── agent.go
    │   ├── application.go
    │   ├── grpc_server.go
    │   ├── http_server.go
    │   ├── kafka.go
    │   ├── logger.go
    │   ├── mertic.go
    │   └── tracing.go
    ├── app
    │   ├── app.go
    │   ├── command
    │   │   ├── create_user.go
    │   │   └── create_user_test.go
    │   └── query
    │       ├── get_profile.go
    │       ├── get_profile_test.go
    │       └── get_token_user.go
    ├── cmd
    │   └── user
    │       └── main.go
    ├── domain
    │   ├── auth.go
    │   ├── mock
    │   │   └── repository.go
    │   ├── repository.go
    │   ├── user.go
    │   └── user_test.go
    ├── metrics
    │   └── metrics.go
    ├── ports
    │   ├── grpc
    │   │   ├── grpc_handlers.go
    │   │   ├── grpc_server.go
    │   │   └── grpc_server_test.go
    │   ├── http
    │   │   ├── handlers.go
    │   │   └── http_server.go
    │   └── kafka
    │       ├── config.go
    │       ├── consumer_group.go
    │       ├── create_user_consumer.go
    │       └── utils.go
    └── proto
        ├── grpc
        │   ├── users_grpc.pb.go
        │   ├── users.pb.go
        │   └── users.proto
        └── kafka
            ├── kafka.pb.go
            └── kafka.proto

```
