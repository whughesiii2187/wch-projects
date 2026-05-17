# Odin Dev Container

Odin development environment with devcontainer CLI + Podman.

## Quick start

```bash
make build      # build the image
make up         # start the container
make shell      # open shell inside
```

## Language server

Copy `ols.json` into your project root to activate OLS (Odin Language Server). Template:

```json
{
  "$schema": "https://raw.githubusercontent.com/DanielGavin/ols/master/misc/ols.schema.json",
  "collections": [
    { "name": "core",   "path": "/opt/odin/core" },
    { "name": "vendor", "path": "/opt/odin/vendor" }
  ],
  "odin_command": "/usr/local/bin/odin"
}
```

## Postgres

If you're using the shared Postgres, make sure it's running:

```bash
cd ../postgres && make up
```

Postgres will be available at `postgres://dev:dev@localhost:5432/devdb?sslmode=disable` inside the container (via `--network=host`).

## What's inside

- Odin (latest nightly)
- `clang`, `llvm`, `lld` (linker toolchain)
- `zsh`, `tmux`, `ripgrep`, `fd`, `fzf`
- `psql` (Postgres client)

Your editor and LSP wiring comes from your dotfiles.
