# backend-playground

This repo is a playground for backend development.

## Summary

- `docker compose up -d` to start the Postgres and Adminer containers
- `./setup.sh` to install tools
- `run dev` to start the server and watch for changes
- `git commit-gen` to generate a commit message based on the changes staged for commit

## Devcontainer

- Ubuntu 22.04, Postgres 15, Golang 1.21, Docker 24.0.9-1, Docker Compose 2.25.0-1

### Ent
Ent is an entity framework of the Go language developed by Facebook. The aim  of ent is to simplify the process of building and maintaining applications with large and complex data models.

Ent provides a functional API to interact with the database, allowing developers to model any database schema as a Graph of Go structs and automatically generate CRUD (Create, Read, Update, Delete) operations.

[Getting Started](https://entgo.io/docs/getting-started)

#### Operations

- `go run -mod=mod entgo.io/ent/cmd/ent new <MODEL_NAME>` to create a new schema entry
- `go generate ./ent` to generate the ent code

## Custom Git Commit Message Hook                                             
                                                                              
This repository includes a custom Git hook designed to enhance the commit process by automatically generating commit messages based on the changes staged for commit, with a focus on quality and conciseness. Below are the details on how to set up and utilize this hook.

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
