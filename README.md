# backend-playground

This repo is a playground for backend development. 

## Environment

- Ubuntu 22.04, Postgres 15, Golang 1.21, Docker 24.0.9-1, Docker Compose 2.25.0-1
- See `.devcontainer` for environment setup
- Postgres and Adminer are started with `docker-compose up -d` or as part of the devcontainer setup

## Development

- `./setup.sh` to install some tools
- `run --list` to see all available commands
- `run dev` to start the server and watch for changes

## Ideas / TODO

Use `ent` to manage the data model and database schema.

### Ent
Ent is an entity framework of the Go language developed by Facebook. The aim  of ent is to simplify the process of building and maintaining applications with large and complex data models.

Ent provides a functional API to interact with the database, allowing developers to model any database schema as a Graph of Go structs and automatically generate CRUD (Create, Read, Update, Delete) operations.

[Getting Started](https://entgo.io/docs/getting-started)

#### Operations

- `go run -mod=mod entgo.io/ent/cmd/ent new <MODEL_NAME>` to create a new schema entry
- `go generate ./ent` to generate the ent code
