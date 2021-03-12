# gin-ca-boilerplate

Go Gin boilerplate, which using clean architecture

## Project structure

- `cmd` - app entrypoint scripts
    - `app` - app instance
    - `consumer` - consumer to work with AMQP (default: Rabbit) queue <em>Optional.</em>
    - `migrations` - app db migrations
- `internal` - project files not intended for external usage. Main features places here.
    - `app`
        - `app.go` - app init entrypoint
    - `config`
        - `config.go` - contain struct with all project config
    - `delivery` - transport layer. Can container many transports (i.g HTTP, GRPC, AMQP)
        - `http` - default app transport
            - `templates` - static template routes
            - `v1` - api version handlers
    - `entity` - service entities structs and interface
    - `enum`
    - `model`
    - `repository` - data source layer. Can container repo implementations for many sources (default: Postgres)
    - `service` - business logic layer
- `pkg` - project files, which can usage from other projects
- `static` - static files. This folder shared via server
- `templates` - .gohtml templates
    

Generated:
- `bin` - executable files
- `docs` - swagger doc files

## Develop

1. Generate Swagger doc (optionally):
```shell
swag init -g ./cmd/app/main.go --parseDependency --parseInternal
```

2. Run server
```shell
go build -o ./bin/app ./cmd/app/main.go

source ./bin/app
```


## Deploy

1. Run:
```shell
docker-compose up
```

2. Destroy
```shell
docker-compose down --rmi 'local'
```
