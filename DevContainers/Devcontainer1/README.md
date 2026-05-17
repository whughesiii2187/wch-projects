# Dev Container

A reproducible, language-agnostic development container running inside Podman via the devcontainer CLI.

Ships with **Go** and **Odin** toolchains. Your editor (Neovim, VSCode, etc.) and its configuration comes from your dotfiles. This container provides only the runtime tooling and binaries you need.

## Prerequisites

| Tool | Install |
|------|---------|
| Podman (rootless) | `dnf install podman` / `apt install podman` |
| devcontainer CLI | `npm i -g @devcontainers/cli` |
| Node ≥ 18 | (for devcontainer CLI) |

```bash
podman info --format '{{.Host.Security.Rootless}}'   # should print true
```

---

## Quick start

```bash
make build      # build the image
make pod-up     # start devcontainer + Postgres sidecar
make pod-shell  # open a shell inside the container
```

---

## Adding a language

1. **Dockerfile** — uncomment or add a language section (Python, Rust stubs are provided)
2. Rebuild: `make rebuild`

That's it. No editor config to maintain.

---

## What's inside

### Base
- **Image**: `debian:bookworm-slim`
- **Shell**: `zsh`, `tmux`
- **Build tools**: `make`, `gcc`, `g++`, `pkg-config`
- **Search**: `ripgrep`, `fd`, `fzf`
- **Database**: `psql` (Postgres runs as a sidecar)

### Go
- Compiler (1.23)
- `dlv` (debugger)
- `goimports`, `gofumpt` (formatting)
- `golangci-lint` (linting)
- `gomodifytags`, `impl` (code generation)

### Odin
- Compiler (latest nightly)
- Linker: `clang`, `llvm`, `lld`

---

## Postgres sidecar

Runs in the same pod at `localhost:5432`.

| Setting | Value |
|---------|-------|
| User | `dev` |
| Password | `dev` |
| Database | `devdb` |
| `DATABASE_URL` | `postgres://dev:dev@localhost:5432/devdb?sslmode=disable` |

Data persists across pod restarts.

---

## Makefile targets

```
make build      — build the container image
make pod-up     — start the pod (devcontainer + Postgres)
make pod-down   — stop and remove the pod
make pod-shell  — open zsh inside the devcontainer
make pod-nvim   — open neovim inside the devcontainer
make pg-cli     — connect to Postgres from the host
make up         — start via devcontainer CLI (no sidecar)
make shell      — shell via devcontainer CLI
make rebuild    — rebuild image without cache
make prune      — remove dangling images and volumes
```
