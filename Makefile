BINARY_NAME=main.out

install_dependencies:
	go mod tidy

build:
	go build -o ${BINARY_NAME} ./cmd/main.go

run:
	go build -o ${BINARY_NAME} ./cmd/main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ./${BINARY_NAME}

test:
	go test -v ./...

create_local_db:
	docker run -d --name data-persistence -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -e POSTGRES_DB=data-persistence postgres:15

migrate_local_db:
	migrate -path "./database/migrations" -database "postgres://root:root@localhost:5432/data-persistence?sslmode=disable" up

compose:
	docker compose up
