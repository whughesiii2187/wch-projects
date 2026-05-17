FROM golang:1.22

# System updates + essential tools
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        git \
        build-essential \
        pkg-config \
        netcat-openbsd \
        iputils-ping \
        postgresql-client \
        && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Optional: create a non-root user
RUN useradd -m devuser
USER devuser
WORKDIR /workspace

