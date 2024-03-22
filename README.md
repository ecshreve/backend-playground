# backend-playground

This repo is a playground for backend development.

## Summary

- VSCode connected to remote host via SSH and the repository is opened in a Docker container for development
  - Ubuntu 22.04 [devcontainer](https://code.visualstudio.com/docs/devcontainers/containers) with Golang 1.21.8.
  - Host docker socket accessible from the devcontainer via the docker-from-docker feature.
- `User` and `Todo` models defined in `ent/schema`, basic CRUD operations generated.
- Database migrations via Ent [automatic migration](https://entgo.io/docs/versioned/intro#automatic-migration).

## Todo
- [ ] Generate a gRPC service with ent
- [ ] Add a `ProfileImage` field on the `User` model
    - [ ] Spin up LocalStack S3 for image storage
- [ ] Enable GraphQL mutations for User and Todo, write tests
- [ ] Look at versioned migrations
- [ ] Generate a REST API with ent
- [ ] Audit tasks.toml and update the README
- [ ] Think about app configuration, possibly move to a TOML or YAML file
- [ ] Reimplement commit generator func into a standalone shell script
	- Remove the script from .githooks, reset git config to default, add configuration option
		or run task to show setting it as a hook.
	- Reimplement the script in a standalone shell script, maybe add to PATH?
	- Add run task for the fancy gum output confirmation
	- Add alias `gcg` to run the script wherever it ends up
- [ ] Figure out a way to log / monitor the graphql requests better.
- [ ] Add custom slog handler with formatting and clors, I have this somewhere just need to find it.
- [ ] Maybe try an AI static site generator to spin up a quick UI
- [?] Iterate on prebuilding the devcontainer image, what's the best way to do that?
---
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


## Custom Git Commit Message Hook                                             
                                                                              
This repository includes a custom Git hook designed to enhance the commit process by automatically generating commit messages based on the changes staged for commit, with a focus on quality and conciseness. Below are the details on how to set up and utilize this hook. `mods-config.yml` is an example `mods` configuration file.

### Setup

Set the .githooks directory as the core.hooksPath and create an alias for the command that calls the hook (`git commit-gen`).
```
git config core.hooksPath .githooks
git config alias.commit-gen '!f() { sh .githooks/generate-commit-msg.sh "$@" && git commit; }; f'
```

### Usage

To use the custom hook, stage the changes you want to commit and run `git commit-gen`. This will generate a commit message based on the changes staged for commit and then continue with the commit process.

The generation script interacts with the OpenAI API via the `mods` CLI tool. The `comm` role in my `mods` configuration uses the following system prompt:

```
"You are a helpful assistant tasked with writing git commit messages. You will be provided a set of code changes to analyze and summarize. You focus on the intent and impact of changes. You output git commit messages in the format <TYPE>(<SCOPE>): <DESCRIPTION>

Generate a concise git commit message written in the active voice and present tense for the given set of changes by following these steps:

Step 1: Summarize the set of changes as a bulleted list and do not output
Step 2: Choose a TYPE from the following list that best describes the change:
	- chore: Other changes that dont modify src or test files
	- docs: Updates documentation
	- feat: A new feature
	- fix: A bug fix
	- test: Adding missing tests or correcting existing tests
Step 3: Identify the SCOPE of the changes with the following constraints:
	- must be a single word describing the area of the repository most affected
	- could be a package name, a directory, or a single filename
Step 4: Generate a concise DESCRIPTION of the changes with the following constraints:
	- must contain 32 characters or less
	- must start with a lowercase letter
	- must end in an alphanumeric character
  - must be written in the active voice and present tense
Step 5: Generate a BODY for the commit message with the following constraints:
  - must contain a bullet list with details of the changes using '-' as the bullet character
  - must not contain more than 72 characters per line
  - must be written in the active voice and present tense
Step 6: Output the git commit message in the format: <TYPE>(<SCOPE>): <DESCRIPTION>\n\n<BODY>"
```
