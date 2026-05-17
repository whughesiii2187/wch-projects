# syntax=docker/dockerfile:1
FROM debian:bookworm-slim

ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=1000

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl wget git unzip tar xz-utils \
    bash zsh tmux \
    make gcc g++ pkg-config \
    ripgrep fd-find fzf \
    jq tree ca-certificates \
    clang llvm lld \
    postgresql-client \
    && rm -rf /var/lib/apt/lists/*

# ── Odin ──────────────────────────────────────────────────────────────────────
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

RUN groupadd --gid ${USER_GID} ${USERNAME} && \
    useradd --uid ${USER_UID} --gid ${USER_GID} -m -s /bin/zsh ${USERNAME}

USER ${USERNAME}
WORKDIR /home/${USERNAME}

ENV ODIN_ROOT=/opt/odin
ENV PATH=/opt/odin:/usr/local/go/bin:${PATH}
ENV DATABASE_URL=postgres://dev:dev@localhost:5432/devdb?sslmode=disable

RUN echo 'export ODIN_ROOT="/opt/odin"' >> /home/${USERNAME}/.zshrc && \
    echo 'export PATH="/opt/odin:$PATH"' >> /home/${USERNAME}/.zshrc

WORKDIR /workspace

CMD [ "zsh" ]
