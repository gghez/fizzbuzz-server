# Demonstration project

Go mod barely used here, mostly for IDE configuration (Goland). May be useful when a project grows.

Run the project with `go run cmd/main.go`

## Project structure

**cmd/** : HTTP server entrypoint

**pkg/fizzbuzz/** : Algorithm related code

**pkg/server/** : HTTP server related code

## Tests

```
go test ./...
```

## Linter

Install golint and then run:

```
golint ./...
```

