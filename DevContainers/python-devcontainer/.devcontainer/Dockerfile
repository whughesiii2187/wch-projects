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
    postgresql-client \
    python3 python3-pip python3-venv \
    && rm -rf /var/lib/apt/lists/*

RUN groupadd --gid ${USER_GID} ${USERNAME} && \
    useradd --uid ${USER_UID} --gid ${USER_GID} -m -s /bin/zsh ${USERNAME}

USER ${USERNAME}
WORKDIR /home/${USERNAME}

ENV PATH=/home/${USERNAME}/.local/bin:/home/${USERNAME}/.venv/bin:${PATH}
ENV DATABASE_URL=postgres://dev:dev@localhost:5432/devdb?sslmode=disable

RUN python3 -m venv /home/${USERNAME}/.venv && \
    /home/${USERNAME}/.venv/bin/pip install --upgrade pip setuptools wheel

RUN echo 'export PATH="$HOME/.local/bin:$HOME/.venv/bin:$PATH"' >> /home/${USERNAME}/.zshrc

WORKDIR /workspace

CMD [ "zsh" ]
