FROM golang:1.20-alpine AS build
WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /build ./cmd/main.go

FROM build AS run-test-stage
RUN go test -v ./...

FROM alpine:3.17 AS release
WORKDIR /newsletter-platform

COPY config.json .
COPY --from=build /build ./build

EXPOSE 8443

ENTRYPOINT [ "/newsletter-platform/build" ]