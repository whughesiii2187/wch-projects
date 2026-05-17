# Makefile — Postgres standalone

.PHONY: up down logs cli prune

## Start Postgres
up:
	podman-compose up -d

## Stop Postgres
down:
	podman-compose down

## Tail logs
logs:
	podman-compose logs -f postgres

## Connect via psql
cli:
	psql postgres://dev:dev@localhost:5432/devdb

## Clean up volumes
prune:
	podman volume prune -f
