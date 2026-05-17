# Python Dev Container

Python development environment with devcontainer CLI + Podman.

## Quick start

```bash
make build      # build the image
make up         # start the container
make shell      # open shell inside
```

## Dependencies

Create a `requirements.txt` in your project and install:

```bash
pip install -r requirements.txt
```

The container has a pre-created venv at `~/.venv` and is activated in `.zshrc`.

## Postgres

If you're using the shared Postgres, make sure it's running:

```bash
cd ../postgres && make up
```

Postgres will be available at `postgres://dev:dev@localhost:5432/devdb?sslmode=disable` inside the container (via `--network=host`).

## What's inside

- Python 3 + pip
- venv at `~/.venv` (pre-activated)
- `zsh`, `tmux`, `ripgrep`, `fd`, `fzf`
- `psql` (Postgres client)

Your editor and LSP wiring comes from your dotfiles.
