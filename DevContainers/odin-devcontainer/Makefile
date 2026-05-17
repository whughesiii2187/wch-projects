# Makefile — Odin devcontainer

WORKSPACE := $(shell pwd)

.PHONY: build shell up down rebuild

## Build the image
build:
	devcontainer build \
	  --workspace-folder $(WORKSPACE) \
	  --docker-path $(shell which podman)

## Open shell in running container
shell:
	devcontainer exec \
	  --workspace-folder $(WORKSPACE) \
	  --docker-path $(shell which podman) \
	  zsh

## Start the devcontainer
up:
	devcontainer up \
	  --workspace-folder $(WORKSPACE) \
	  --docker-path $(shell which podman)

## Stop the devcontainer
down:
	podman ps -a --format '{{.Names}}' | grep devcontainer | xargs -r podman rm -f

## Rebuild without cache
rebuild:
	devcontainer build \
	  --workspace-folder $(WORKSPACE) \
	  --docker-path $(shell which podman) \
	  --no-cache
