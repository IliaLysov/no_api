FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

FROM alpine:3.20 AS run

COPY --from=build /app/main /main

EXPOSE 8080
ENTRYPOINT ["/main"]