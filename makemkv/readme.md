makemkv builder for ubuntu

Build debs
* `DOCKER_BUILDKIT=1 docker build --target export --output . .` build deb packages and saves in same directory

install debs and run makemkv
* `sudo dpkg -i makemkv-bin_*_amd64.deb makemkv-oss_*_amd64.deb` (install the debs
* `makemkvcon reg keyfromforum` (register the key from forum)
* `makemkv` (run gui)
* `makemkvcon` (run cli)

run in docker
* `docker compose up -d --build` (start container)
* `docker exec -it makemkv makemkvcon reg keyfromforum` (register the key)
* `docker exec -it makemkv makemkvcon`(run from container)

update version
note: if you want to compile a newer version change these 
* change `FROM ubuntu:25.10 AS builder` and `ENV MAKEMKV_VERSION=1.18.2` in `Dockerfile` and then change `image: makemkv:1.18.2` in `docker-compose.yaml` to same makemkv version
