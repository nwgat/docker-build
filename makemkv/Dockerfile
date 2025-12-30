# syntax=docker/dockerfile:1

# --- Stage 1: Builder ---
FROM ubuntu:25.10 AS builder

ENV DEBIAN_FRONTEND=noninteractive
ENV MAKEMKV_VERSION=1.18.2

# Checksums
ENV OSS_SHA256=b9497f5555c257f5d3c7f18b6ab03c74cff3ea4437e80349d618d115bf5fa445
ENV BIN_SHA256=bfc4c7cebc0f00497671ffd56e6a34f0770a9e6af7eff2f0127efa143f28636e

# Install build tools
RUN apt-get update && apt-get install -y \
    build-essential pkg-config libc6-dev libssl-dev libexpat1-dev \
    libavcodec-dev libgl1-mesa-dev qtbase5-dev zlib1g-dev \
    wget less checkinstall \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /tmp/makemkv-build

# 1. Build OSS .deb
RUN wget https://www.makemkv.com/download/makemkv-oss-${MAKEMKV_VERSION}.tar.gz && \
    echo "${OSS_SHA256}  makemkv-oss-${MAKEMKV_VERSION}.tar.gz" | sha256sum -c - && \
    tar -xvzf makemkv-oss-${MAKEMKV_VERSION}.tar.gz && \
    cd makemkv-oss-${MAKEMKV_VERSION} && \
    ./configure && \
    checkinstall --default --install=no --pkgname=makemkv-oss --pkgversion=${MAKEMKV_VERSION} --fstrans=no make install && \
    mkdir -p /packages && \
    mv *.deb /packages/

# 2. Build BIN .deb
RUN wget https://www.makemkv.com/download/makemkv-bin-${MAKEMKV_VERSION}.tar.gz && \
    echo "${BIN_SHA256}  makemkv-bin-${MAKEMKV_VERSION}.tar.gz" | sha256sum -c - && \
    tar -xvzf makemkv-bin-${MAKEMKV_VERSION}.tar.gz && \
    cd makemkv-bin-${MAKEMKV_VERSION} && \
    mkdir -p tmp && touch tmp/eula_accepted && \
    checkinstall --default --install=no --pkgname=makemkv-bin --pkgversion=${MAKEMKV_VERSION} --fstrans=no make install && \
    mv *.deb /packages/

# --- Stage 2: Export (For getting .deb files to host) ---
# This stage is used when you run: docker build --target export --output . .
FROM scratch AS export
COPY --from=builder /packages/ .

# --- Stage 3: Runtime (Default for Docker Compose) ---
FROM ubuntu:25.10 AS runtime
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && apt-get install -y \
    libssl3 libexpat1 libavcodec61 libgl1 libqt5widgets5 zlib1g less \
    && rm -rf /var/lib/apt/lists/*

# Copy .deb files from builder
COPY --from=builder /packages/*.deb /tmp/

# Install them
RUN dpkg -i /tmp/makemkv-oss_*.deb /tmp/makemkv-bin_*.deb && \
    rm /tmp/*.deb

# Runtime Config
RUN mkdir -p /output   
WORKDIR /output
CMD ["makemkvcon"]
