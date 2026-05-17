# syntax=docker/dockerfile:1
# ── Base ──────────────────────────────────────────────────────────────────────
FROM debian:bookworm-slim

ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=1000

# ── System packages ───────────────────────────────────────────────────────────
RUN apt-get update && apt-get install -y --no-install-recommends \
    # Core utilities
    curl \
    wget \
    git \
    unzip \
    tar \
    xz-utils \
    # Shell & terminal
    bash \
    zsh \
    tmux \
    # Build essentials
    make \
    gcc \
    g++ \
    pkg-config \
    # Search & navigation
    ripgrep \
    fd-find \
    fzf \
    # General utilities
    jq \
    tree \
    ca-certificates \
    # Clang — required linker for Odin on Linux
    clang \
    llvm \
    lld \
    # Postgres client — server runs as a sidecar (see pod.yaml)
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

# ── Go ────────────────────────────────────────────────────────────────────────
# To remove: delete this section and the Go tools block below.
ARG GO_VERSION=1.23.0
RUN curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
      -o /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

ENV PATH=/usr/local/go/bin:${PATH}

# CLI tools only — LSP (gopls) and editor plugins are managed by your dotfiles/Neovim
RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
    go install mvdan.cc/gofumpt@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest

# ── Odin ──────────────────────────────────────────────────────────────────────
# To remove: delete this section. Also remove clang/llvm/lld above if no other language needs them.
RUN ODIN_URL=$(curl -s https://api.github.com/repos/odin-lang/Odin/releases/latest \
      | grep '"browser_download_url"' \
      | grep 'linux.*amd64' \
      | head -1 \
      | cut -d'"' -f4) && \
    mkdir -p /opt/odin && \
    curl -fsSL "$ODIN_URL" -o /tmp/odin.tar.gz && \
    tar -C /opt/odin -xzf /tmp/odin.tar.gz --strip-components=1 && \
    rm /tmp/odin.tar.gz && \
    ln -sf /opt/odin/odin /usr/local/bin/odin

# ── Add a language: Python (uncomment to enable) ──────────────────────────────
# RUN apt-get update && apt-get install -y --no-install-recommends \
#     python3 python3-pip python3-venv \
#     && rm -rf /var/lib/apt/lists/*

# ── Add a language: Rust (uncomment to enable) ────────────────────────────────
# ENV CARGO_HOME=/usr/local/cargo
# ENV RUSTUP_HOME=/usr/local/rustup
# ENV PATH=/usr/local/cargo/bin:${PATH}
# RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs \
#       | sh -s -- -y --no-modify-path --profile minimal

# ── Non-root user (mirrors host UID/GID for Podman rootless) ──────────────────
RUN groupadd --gid ${USER_GID} ${USERNAME} && \
    useradd --uid ${USER_UID} --gid ${USER_GID} -m -s /bin/zsh ${USERNAME}

USER ${USERNAME}
WORKDIR /home/${USERNAME}

RUN mkdir -p /home/${USERNAME}/go/{bin,src,pkg}

ENV GOPATH=/home/${USERNAME}/go
ENV GOBIN=/home/${USERNAME}/go/bin
ENV ODIN_ROOT=/opt/odin
ENV PATH=/home/${USERNAME}/go/bin:/usr/local/go/bin:/opt/odin:${PATH}

RUN echo 'export GOPATH="$HOME/go"'                                      >> /home/${USERNAME}/.zshrc && \
    echo 'export ODIN_ROOT="/opt/odin"'                                  >> /home/${USERNAME}/.zshrc && \
    echo 'export PATH="$HOME/go/bin:/usr/local/go/bin:/opt/odin:$PATH"' >> /home/${USERNAME}/.zshrc

WORKDIR /workspace

CMD [ "zsh" ]
