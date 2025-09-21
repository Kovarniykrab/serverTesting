build:
    rm -f ./bin/srv
	go build -o ./bin/srv ./cmd/main.go
swagger:
    swag init --parceDependency -g cmd/main.go

lint:
	gofmt -s -w ./ && golangci-lint run

migrate-up:
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgresql goose -dir=./resources/store/psql/migrations up
migrate-down:
	GOOSE_DRIVER=postgres goose -dir=./resources/store/psql/migrations down
migrate-create:
	GOOSE_DRIVER=postgres goose -dir=./resources/store/psql/migrations create new psql