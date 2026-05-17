# Makefile — convenience wrappers around devcontainer CLI + Podman
# Requires: devcontainer CLI  (npm i -g @devcontainers/cli)
#           Podman             (rootless)

CONTAINER_NAME ?= devcontainer
WORKSPACE      := $(shell pwd)

.PHONY: up down shell build rebuild logs prune pod-up pod-down pod-shell pg-cli

## Build image and start the devcontainer
up:
	devcontainer up \
	  --workspace-folder $(WORKSPACE) \
	  --docker-path $(shell which podman) \
	  --id-label devcontainer.local_folder=$(WORKSPACE)

## Open an interactive shell inside the running container
shell:
	devcontainer exec \
	  --workspace-folder $(WORKSPACE) \
	  --docker-path $(shell which podman) \
	  zsh

## Open nvim directly inside the container
nvim:
	devcontainer exec \
	  --workspace-folder $(WORKSPACE) \
	  --docker-path $(shell which podman) \
	  nvim .

## Stop and remove the container
down:
	podman ps -a --format '{{.Names}}' | grep $(CONTAINER_NAME) | \
	  xargs -r podman rm -f

## Rebuild image from scratch (no cache)
rebuild:
	devcontainer build \
	  --workspace-folder $(WORKSPACE) \
	  --docker-path $(shell which podman) \
	  --no-cache

## Tail container logs
logs:
	podman logs -f $(CONTAINER_NAME)

## Remove dangling images / build cache
prune:
	podman image prune -f
	podman system prune -f --volumes

# ── Podman pod targets (Postgres sidecar) ─────────────────────────────────────

## Build the devcontainer image (required before pod-up)
build:
	devcontainer build \
	  --workspace-folder $(WORKSPACE) \
	  --docker-path $(shell which podman)

## Start the full pod (devcontainer + Postgres sidecar)
pod-up:
	podman kube play pod.yaml

## Stop and tear down the pod
pod-down:
	podman kube down pod.yaml

## Open a shell in the devcontainer inside the pod
pod-shell:
	podman exec -it dev-pod-devcontainer zsh

## Open nvim in the devcontainer inside the pod
pod-nvim:
	podman exec -it dev-pod-devcontainer nvim .

## Connect to Postgres via psql (from the host)
pg-cli:
	psql postgres://dev:dev@localhost:5432/devdb
