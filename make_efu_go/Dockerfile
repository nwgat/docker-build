# run DOCKER_BUILDKIT=1 docker build --output . .
# to build this

# 1. Build Stage
FROM golang:1.25 AS builder

WORKDIR /app

# Initialize module and copy source
COPY make_efu.go .
RUN go mod init make_efu

# Compile for Linux AMD64
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o make_efu make_efu.go

# Compile for Linux ARMv7 (armhf)
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=7
RUN go build -o make_efu_arm make_efu.go


# 2. Output Stage (Empty scratch image)
FROM scratch AS export
COPY --from=builder /app/make_efu /
COPY --from=builder /app/make_efu_arm /
