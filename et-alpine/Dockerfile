# syntax=docker/dockerfile:1

FROM alpine:3.23.2 AS builder

# 1. Setup Alpine Build System (abuild)
RUN sed -i '/community/s/^#//' /etc/apk/repositories && \
    apk update && \
    apk add --no-cache alpine-sdk cmake sudo \
    openssl-dev protobuf-dev libsodium-dev curl-dev \
    zlib-dev xz-dev libselinux-dev libutempter-dev libunwind-dev \
    linux-headers

# 2. Create 'builder' user
RUN adduser -D builder && \
    addgroup builder abuild && \
    echo "builder ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers

# 3. Switch to 'builder', Generate Keys, and Trust Them
USER builder
# Generate keys and copy public key to system keys so the indexer trusts it
RUN abuild-keygen -a -n && \
    sudo cp /home/builder/.abuild/*.rsa.pub /etc/apk/keys/

# 4. Create the APKBUILD recipe
WORKDIR /home/builder/package
COPY <<'EOF' APKBUILD
# Contributor: Gemini <ai@google.com>
# Maintainer: You <you@example.com>
pkgname=eternalterminal
pkgver=6.2.4
pkgrel=0
pkgdesc="Remote shell that automatically reconnects without interrupting the session"
url="https://mistertea.github.io/EternalTerminal/"
arch="x86_64"
license="Apache-2.0"
makedepends="cmake build-base openssl-dev protobuf-dev libsodium-dev curl-dev zlib-dev xz-dev git libselinux-dev libutempter-dev libunwind-dev linux-headers"
subpackages="et:client et-server:server"

source=""
options="!check"

builddir="$srcdir/$pkgname"

prepare() {
    mkdir -p "$srcdir"
    cd "$srcdir"

    if [ ! -d "$pkgname" ]; then
        git clone --recurse-submodules --depth 1 https://github.com/MisterTea/EternalTerminal.git "$pkgname"
    fi
    
    cd "$pkgname"

    sed -i \
        -e 's/add_subdirectory(${EXTERNAL_DIR}\/Catch2)/#&/' \
        -e 's/find_package(Catch2/#&/' \
        -e '/\banl\b/d' \
        -e '/Catch2::Catch2WithMain/d' \
        CMakeLists.txt
}

build() {
    mkdir -p build && cd build
    cmake .. \
        -G "Unix Makefiles" \
        -DCMAKE_INSTALL_PREFIX=/usr \
        -DBUILD_TESTING=OFF \
        -DDISABLE_SENTRY=ON \
        -DDISABLE_TELEMETRY=ON \
        -DDISABLE_VCPKG=ON \
        -DLIBUNWIND_ENABLE_SHARED=OFF \
        -DLIBUNWIND_ENABLE_STATIC=OFF
    make -j$(nproc)
}

package() {
    cd "$builddir/build"
    make install DESTDIR="$pkgdir"
}

client() {
    pkgdesc="Eternal Terminal Client"
    depends="libsodium protobuf libcurl libunwind libutempter"
    mkdir -p "$subpkgdir"/usr/bin
    mv "$pkgdir"/usr/bin/et "$subpkgdir"/usr/bin/
}

server() {
    pkgdesc="Eternal Terminal Server"
    depends="libsodium protobuf libcurl libunwind libutempter"
    mkdir -p "$subpkgdir"/usr/bin
    mv "$pkgdir"/usr/bin/etserver "$subpkgdir"/usr/bin/
    mv "$pkgdir"/usr/bin/etterminal "$subpkgdir"/usr/bin/
}
EOF

# 5. Run the Build
RUN abuild -r

# --- Stage 2: Export ---
FROM scratch AS export
# FIX: Use wildcard '*' for the repo name to ensure we catch the folder
# The logs showed the folder is named 'builder' (e.g. /home/builder/packages/builder/x86_64)
COPY --from=builder /home/builder/packages/*/x86_64/*.apk /
