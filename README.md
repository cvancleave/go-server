# go-server

This application is an basic implementation an API server in Go.
It includes:
- JWT auth and validation.
- Terraform for infrastructure.
- Github actions for deployments.
- Dockerfile for multi-step server build.
- SQL database integration (soon!)
- Flyway database migration (soon!)

Uses Go version `1.21` and the following packages:
- `github.com/aws/aws-sdk-go-v2` for storing secret values.
- `github.com/golang-jwt/jwt` for encoding/decoding tokens.
- `github.com/julienschmidt/httprouter` for the server http router.

## Running the server

Get config and start server.
- Setup server config in AWS secretsmanager.
- Run `go run cmd/server/main.go`

## Running the client

Get config and start client, which creates a token and then hits the server endpoints.
- Make sure server config is setup in AWS secretsmanager.
- Run `go run cmd/client/main.go`

