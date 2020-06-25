# Simple GO HTTP Server

## Unit test

```sh
go test ./app -v -count=1
```

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

## Install GO on Ubuntu 18.04

```sh
wget https://dl.google.com/go/go1.14.linux-amd64.tar.gz
```

```sh
sudo tar -xvf go1.14.linux-amd64.tar.gz
sudo mv go /usr/local
```

```sh
vi ~/.bashrc

# Add to end of file
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

source ~/.bashrc
```

## Docker

### Build docker image

```sh
docker build -t simple-go-http-server .
```

### Migrate database by docker

```sh
docker run --rm -v $(pwd)/config.yaml:/root/config.yaml simple-go-http-server:latest migrate-db
```

### Run HTTP server by docker

```sh
docker run -d -p 8888:8888 -v $(pwd)/config.yaml:/root/config.yaml simple-go-http-server:latest serve
```
