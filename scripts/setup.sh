#!/bin/sh
# Temporary spot for environment setup that is not covered by devcontainer.

# Install go tools
go install github.com/amonks/run/cmd/run@latest
go install github.com/charmbracelet/mods@latest
go install github.com/charmbracelet/glow@latest
go install github.com/charmbracelet/gum@latest

# Install protoc and proto generation tools
apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@master

# Install atlasgo for managing go db migrations
curl -sSf https://atlasgo.sh | sh