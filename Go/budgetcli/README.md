# Go Dev Container

Go development environment with devcontainer CLI + Podman.

## Quick start

```bash
make build      # build the image
make up         # start the container
make shell      # open shell inside
```

## Postgres

If you're using the shared Postgres, make sure it's running:

```bash
cd ../postgres && make up
```

Postgres will be available at `postgres://dev:dev@localhost:5432/devdb?sslmode=disable` inside the container (via `--network=host`).

## What's inside

- Go 1.23
- `dlv`, `goimports`, `gofumpt`, `golangci-lint`, `gomodifytags`, `impl`
- `zsh`, `tmux`, `ripgrep`, `fd`, `fzf`
- `psql` (Postgres client)

Your editor and LSP wiring comes from your dotfiles.
