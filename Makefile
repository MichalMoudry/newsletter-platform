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

# TODO: Add docker comands