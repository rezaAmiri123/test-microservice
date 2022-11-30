Hi everyone<br/> 
This is a test project about implementing a library <br/>
some technologies have been used in this project<br/>

[Kafka](https://github.com/segmentio/kafka-go) as messages broker<br/>
[gRPC](https://github.com/grpc/grpc-go) Go implementation of gRPC<br/>
[PostgreSQL](https://github.com/jackc/pgx) as database<br/>
[Jaeger](https://www.jaegertracing.io/) open source, end-to-end distributed <br/>
[tracing](https://opentracing.io/) tracing<br/>
[Prometheus](https://prometheus.io/) monitoring and alerting<br/>
[Redis](https://github.com/go-redis/redis) Type-safe Redis client for Golang<br/>
[Echo](https://github.com/labstack/echo) web framework<br/>


### To run project on Docker
```bash
make docker_dev
```

### To stop project on Docker
```bash
make docker_dev_down
```

### To run project on kuberneties
```bash
make k8s_install
```

### To stop project on kuberneties
```bash
make k8s_uninstall
```

### Jaeger UI:

http://localhost:16686

### Prometheus UI:

http://localhost:9090


```text
.
├── api_service
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
│   │   ├── command
│   │   │   ├── create_article.go
│   │   │   └── create_user.go
│   │   └── query
│   ├── cmd
│   │   └── api
│   │       └── main.go
│   ├── domain
│   │   └── dto
│   │       ├── create_article.go
│   │       ├── create_user.go
│   │       └── user_login.go
│   ├── metrics
│   │   └── metrics.go
│   └── ports
│       └── http
│           └── v1
│               ├── create_article.go
│               ├── create_user.go
│               ├── http_server.go
│               └── user_loign.go
├── deploy
│   └── test-microservice
│       ├── Chart.yaml
│       ├── templates
│       │   ├── deployment.yaml
│       │   ├── _helpers.tpl
│       │   ├── mysql-deployment.yaml
│       │   ├── mysql-py.yaml
│       │   ├── NOTES.txt
│       │   └── tests
│       │       └── test-connection.yaml
│       └── values.yaml
├── docker
│   ├── api_service.Dockerfile
│   ├── library_service.Dockerfile
│   ├── message_service.Dockerfile
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
│   │       ├── article_repository_test.go
│   │       └── sql_queries.go
│   ├── agent
│   │   ├── agent.go
│   │   ├── application.go
│   │   ├── auth_client.go
│   │   ├── http_server.go
│   │   ├── kafka.go
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
│   ├── ports
│   │   ├── dto
│   │   │   └── article_create.go
│   │   ├── http
│   │   │   ├── create_aricle.go
│   │   │   ├── get_article_by_slug.go
│   │   │   └── http_server.go
│   │   └── kafka
│   │       ├── config.go
│   │       ├── consumer_group.go
│   │       ├── create_article_consumer.go
│   │       └── utils.go
│   └── proto
│       └── kafka
│           ├── kafka.pb.go
│           └── kafka.proto
├── Makefile
├── message_service
│   ├── adapters
│   │   └── pg
│   │       ├── email_repository.go
│   │       └── sql_queries.go
│   ├── agent
│   │   ├── agent.go
│   │   ├── application.go
│   │   ├── kafka.go
│   │   ├── logger.go
│   │   ├── mertic.go
│   │   └── tracing.go
│   ├── app
│   │   ├── app.go
│   │   ├── commands
│   │   │   └── create_email.go
│   │   └── queries
│   │       └── get_email_by_slug.go
│   ├── cmd
│   │   └── message
│   │       └── main.go
│   ├── domain
│   │   └── email
│   │       ├── email.go
│   │       ├── mock
│   │       │   └── repository.go
│   │       └── repository.go
│   ├── metrics
│   │   └── metrics.go
│   ├── ports
│   │   └── kafka
│   │       ├── config.go
│   │       ├── consumer_group.go
│   │       ├── create_email_consumer.go
│   │       └── utils.go
│   └── proto
│       └── kafka
│           ├── kafka.pb.go
│           └── kafka.proto
├── migrations
│   ├── 01_create_initial_tables.down.sql
│   └── 01_create_initial_tables.up.sql
├── monitoring
│   └── prometheus_docker.yml
├── pkg
│   ├── auth
│   │   ├── auth_client
│   │   │   └── grpc.go
│   │   ├── auth.go
│   │   ├── authorize.go
│   │   ├── middleware.go
│   │   └── tls
│   │       ├── files.go
│   │       └── tls.go
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
│   │   ├── mock
│   │   │   └── producer.go
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
│   │   ├── jaeger.go
│   │   └── utils.go
│   └── utils
│       └── validator.go
├── README.md
├── test
│   ├── kafka1
│   │   ├── consumer.go
│   │   └── producer.go
│   ├── main.go
│   └── message_example
│       └── producer.go
├── tls
│   ├── ca-config.json
│   ├── ca-csr.json
│   ├── client-csr.json
│   ├── keys
│   └── server-csr.json
└── user_service
    ├── adapters
    │   ├── user_gorm_repository.go
    │   ├── user_gorm_repository_test.go
    │   └── user_memory_repository.go
    ├── agent
    │   ├── agent.go
    │   ├── application.go
    │   ├── check_alive.go
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
            ├── user_kafka.pb.go
            └── user_kafka.proto
```
