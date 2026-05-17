# Projects

Modular dev containers for multiple languages. Each language is independent; Postgres is shared.

```
projects/
├── postgres/              ← Standalone Postgres, start once
├── go-devcontainer/       ← Go projects
├── odin-devcontainer/     ← Odin projects
└── python-devcontainer/   ← Python projects
```

## Quick start

1. **Start Postgres** (if needed):
   ```bash
   cd postgres && make up
   ```

2. **Build and enter a devcontainer**:
   ```bash
   cd go-devcontainer && make build && make up && make shell
   ```

## How it works

- Each `*-devcontainer` folder is completely self-contained. `cd` into it and run `make build`, `make up`, `make shell`.
- All devcontainers use `--network=host`, so they can reach Postgres at `localhost:5432`.
- `DATABASE_URL` is pre-set to `postgres://dev:dev@localhost:5432/devdb?sslmode=disable` in all containers.
- Your editor (Neovim, VSCode) and LSP wiring comes from your dotfiles — these containers provide only the runtime.

## Adding a project

Create a new folder at the root (`my-go-project/`, `my-odin-project/`, etc.) and copy in whatever language's `.devcontainer/` folder.

Or put multiple projects in a `workspace/` subdirectory and bind that when entering the container.

## Postgres credentials

All containers point to:
- **Host**: `localhost:5432`
- **User**: `dev`
- **Password**: `dev`
- **Database**: `devdb`
