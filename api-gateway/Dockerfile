FROM golang:1.24-alpine AS build
RUN apk add --no-cache git ca-certificates
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY api-gateway ./api-gateway
COPY services/auth-service ./services/auth-service
COPY proto ./proto

RUN go build -o /app/main ./api-gateway/cmd

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE 8080
ENTRYPOINT ["/app/main"]