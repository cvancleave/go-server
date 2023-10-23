# go-server

This application is an basic implementation of an API server in Go.
It includes:
- JWT authentication and validation.
- Terraform and AWS for infrastructure.
- Github actions for deployments.
- Dockerfile for multi-step server build.
- SQL database integration.
- Flyway database migration.

Uses Go version `1.21` and the following packages:
- `github.com/aws/aws-sdk-go-v2` for storing secret values.
- `github.com/golang-jwt/jwt` for encoding/decoding tokens.
- `github.com/julienschmidt/httprouter` for the server http router.

## Running the server

Get config and start server.
- Setup server config in AWS secretsmanager (or dummy config).
- Run `go run cmd/server/main.go`

## Running the client

Get config and start client, which tests the server endpoints.
- Make sure server config is setup in AWS secretsmanager (or dummy config).
- Run `go run cmd/client/main.go`

## Migrating the database

Runs database migrations.
- Setup database, and add to `flyway.conf` to direct flyway.
- Run `cd database` and `flyway migrate`, or `flyway repair` to fix any issues.
