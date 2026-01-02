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
