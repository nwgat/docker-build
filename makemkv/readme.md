makemkv builder for ubuntu

* `DOCKER_BUILDKIT=1 docker build --target export --output . .` build deb packages and saves in same directory
* `docker compose up -d --build` (start container)
* `sudo dpkg -i makemkv-bin_*_amd64.deb makemkv-oss_*_amd64.deb` (install the debs
* `makemkvcon reg keyfromforum` (register the key)
* `makemkv` (run gui)
* `makemkvcon` (run cli)

run in docker

* `docker exec -it makemkv makemkvcon reg keyfromforum` (register the key)
* `docker exec -it makemkv makemkvcon`(run from container)
