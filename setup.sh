#!/bin/sh
# Temporary spot for environment setup that is not covered by devcontainer.

# Install go tools
go install github.com/amonks/run/cmd/run@latest
go install github.com/charmbracelet/mods@latest
go install github.com/charmbracelet/glow@latest