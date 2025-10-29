FROM golang:1.24-alpine AS build
RUN apk add --no-cache git ca-certificates
WORKDIR /app

# use repo-level go.mod so modules are resolvable (build.context must be repo root)
COPY go.mod go.sum ./
RUN go mod download

# copy service sources and proto files into the image
COPY services/auth-service ./services/auth-service
COPY proto ./proto

# build a static binary for the auth service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/auth-service ./services/auth-service/cmd

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=build /usr/local/bin/auth-service /usr/local/bin/auth-service
EXPOSE 50051
ENTRYPOINT ["/usr/local/bin/auth-service"]