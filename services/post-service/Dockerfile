FROM golang:1.24-alpine AS build
RUN apk add --no-cache git ca-certificates
WORKDIR /app

# use repo-level go.mod so modules are resolvable
COPY go.mod go.sum ./
RUN go mod download

# copy service sources and proto files
COPY services/post-service ./services/post-service
COPY proto ./proto

# build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/post-service ./services/post-service/cmd

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=build /usr/local/bin/post-service /usr/local/bin/post-service
EXPOSE 50053
ENTRYPOINT ["/usr/local/bin/post-service"]