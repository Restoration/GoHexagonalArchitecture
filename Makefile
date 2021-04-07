.PHONY: init
init:
	cp ./infrastructure/env/dotenv.local .env

.PHONY: fmt
fmt:
	gofmt -w -s .

.PHONY: lint
lint:
	golint ./...

.PHONY: build
build:
	go build -o bin/main main.go

.PHONY: migrate
migrate:
	sql-migrate up

.PHONY: seed
seed:
	go run ./infrastructure/mysql/seeds/seed.go