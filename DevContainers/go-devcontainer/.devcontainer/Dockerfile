# syntax=docker/dockerfile:1
FROM docker.io/library/debian:bookworm-slim

ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=1000

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl wget git unzip tar xz-utils \
    bash zsh tmux \
    make gcc g++ pkg-config \
    ripgrep fd-find fzf \
    jq tree ca-certificates \
    sudo \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

# ── Go ────────────────────────────────────────────────────────────────────────
ARG GO_VERSION=1.23.0
RUN curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
      -o /tmp/go.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz

ENV PATH=/usr/local/go/bin:${PATH}

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
    go install mvdan.cc/gofumpt@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest

RUN groupadd --gid ${USER_GID} ${USERNAME} && \
    useradd --uid ${USER_UID} --gid ${USER_GID} -m -s /bin/zsh ${USERNAME} && \
    echo "${USERNAME} ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers.d/${USERNAME}

USER ${USERNAME}
WORKDIR /home/${USERNAME}

RUN mkdir -p /home/${USERNAME}/go/{bin,src,pkg}

ENV GOPATH=/home/${USERNAME}/go
ENV GOBIN=/home/${USERNAME}/go/bin
ENV PATH=/home/${USERNAME}/go/bin:/usr/local/go/bin:${PATH}
ENV DATABASE_URL=postgres://dev:dev@localhost:5432/devdb?sslmode=disable

RUN echo 'export GOPATH="$HOME/go"' >> /home/${USERNAME}/.zshrc && \
    echo 'export PATH="$HOME/go/bin:/usr/local/go/bin:$PATH"' >> /home/${USERNAME}/.zshrc

WORKDIR /workspace

CMD [ "zsh" ]
