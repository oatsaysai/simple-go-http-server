# Simple GO HTTP Server

## Run postgres container

```sh
docker run --rm -p 5432:5432 -e POSTGRES_PASSWORD=postgres --name postgres postgres:10-alpine
```

## Create database to Postgres (Drop and Create)

```sh
docker exec -i postgres psql -U postgres -c "drop database if exists simple_db" && \
docker exec -i postgres psql -U postgres -c "create database simple_db"
```

## Migrate database

```sh
go run main.go migrate-db
```

## Run HTTP server

```sh
go run main.go serve
```

## Insert string

```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"data":"test_string_001"}' \
  http://localhost:8888/api/insert
```

## Check string is already exist

```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"data":"test_string_004"}' \
  http://localhost:8888/api/check_exist
```

## List all string in database

```sh
curl \
  --request GET \
  http://localhost:8888/api/list_string
```
