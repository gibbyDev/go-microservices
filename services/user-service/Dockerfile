FROM golang:1.24-alpine AS build
RUN apk add --no-cache git ca-certificates
WORKDIR /app

# use repo-level go.mod so modules are resolvable
COPY go.mod go.sum ./
RUN go mod download

# copy service sources and proto files
COPY services/user-service ./services/user-service
COPY proto ./proto

# build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/user-service ./services/user-service/cmd

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=build /usr/local/bin/user-service /usr/local/bin/user-service
EXPOSE 50052
ENTRYPOINT ["/usr/local/bin/user-service"]