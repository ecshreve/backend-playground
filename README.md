# backend-playground

This repo is a playground for backend development.

## Summary

- VSCode connected to remote host via SSH and the repository is opened in a Docker container for development
  - Ubuntu 22.04 [devcontainer](https://code.visualstudio.com/docs/devcontainers/containers) with Golang 1.21.8.
  - Host docker socket accessible from the devcontainer via the docker-from-docker feature.
- `User` and `Todo` models defined in `ent/schema`, basic CRUD operations generated.
- Database migrations via Ent [automatic migration](https://entgo.io/docs/versioned/intro#automatic-migration).

## Roadmap
- [ ] Audit `dev` run task and record demo
- [ ] Add a `ProfileImage` field on the `User` model
    - [ ] Spin up LocalStack S3 for image storage
- [ ] Enable GraphQL mutations for User and Todo, write tests
- [ ] Look at versioned migrations
- [ ] Generate a REST API with ent
- [ ] Audit tasks.toml and update the README
- [ ] Think about app configuration, possibly move to a TOML or YAML file
- [ ] Figure out a way to log / monitor the graphql requests better.
- [ ] Add custom slog handler with formatting and clors, I have this somewhere just need to find it.
- [ ] Maybe try an AI static site generator to spin up a quick UI
- [?] Iterate on prebuilding the devcontainer image, what's the best way to do that?
---
- [x] Write logs to file and configure Loki for log viewing
- [x] Add logging middleware to grpcserver
- [x] Generate a gRPC service with ent
- [x] Reimplement commit generator func into a standalone shell script
	- [x] Remove the script from .githooks, reset git config to default
	- ~~add configuration option or run task to show setting it as a hook.~~
	- [x] Reimplement the script in a standalone shell script, add scripts folder to PATH
	- ~~Add run task for the fancy gum output confirmation~~
	- ~~Add alias `gcg` to run the script wherever it ends up~~
- [x] Audit tasks.toml and update the README
- [x] Add an edge between User and Todo
- [x] Populate the database with dummy seed data
- [x] Generate a GraphQL API with ent

## Development Commands

- `docker compose up -d` to start the Postgres and Adminer containers
  - or `run dbdev` to start the Postgres and Adminer containers inside `run` tasks
- `./setup.sh` to install tools
- `run dev` to start the server and watch for changes
- `git commit-gen` to generate a commit message based on the changes staged for commit
- `go run -mod=mod entgo.io/ent/cmd/ent new <MODEL_NAME>` to generate a new model

### Devcontainer

- Ubuntu 22.04, Golang 1.21, Docker 24.0.9-1, Docker Compose 2.25.0-1
- Based on the `base:ubuntu-22.04` devcontainer image: `mcr.microsoft.com/devcontainers/base:jammy`

## Ent

Ent is an entity framework of the Go language developed by Facebook. The aim  of ent is to simplify the process of building and maintaining applications with large and complex data models.

Ent provides a functional API to interact with the database, allowing developers to model any database schema as a Graph of Go structs and automatically generate CRUD (Create, Read, Update, Delete) operations.

[Getting Started](https://entgo.io/docs/getting-started)


## GraphQL

Generate a GraphQL API with `run gqlgen` and then `run gqlserver` to start the server.

`ent/entc.go`
- Uses the (entql)[https://pkg.go.dev/entgo.io/contrib/entgql] extension to generate a GraphQL schema and resolvers for the Ent models.

`gqlserver/ent.graphql`
- The generated GraphQL schema file for the Ent models.
- The schema is defined using the GraphQL Schema Definition Language (SDL).

`gqlserver/resolver.go`
- The generated GraphQL base Resolver.

`gqlserver/ent.resolvers.go`
- The generated Ent resolvers that extend the base for the different Ent models.

`gqlserver/gql-generated.go`
- The generated GraphQL types and interfaces needed to implement the resolvers and server. 


## Git Commit Message Script

This repository includes a script designed to enhance the commit process by automatically generating commit messages based on the changes staged for commit, with a focus on quality and conciseness. It uses `mods` to send the current `git diff` in addition to a custom prompt to the OpenAI API. Below are the details on how to set up and use the script. `mods-config.yml` is the config used in the script.

### Setup

- My set up is to add the `scripts/` directory to my PATH and set the OPENAI_API_KEY environment variable in this repo's .env file.
- The script uses `mods` and `gum` CLI tools to interact with the OpenAI API and generate the commit message, they can be installed by running the setup script: `./scripts/setup.sh`.
- See the config file `mods-config.yml` for all the LLM and prompt related settings.

### Usage

To use the script, stage the changes you want to commit and run it: `./scripts/gencom.sh`. 

<img src="assets/demo.gif" width="400" alt="Demo GIF">


### Resources

- [OpenAI API](https://beta.openai.com/docs/)
- [Gum](https://github.com/charmbracelet/gum)
- [Mods](https://github.com/charmbracelet/mods)

