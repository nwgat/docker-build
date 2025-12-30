based on make_efu perl script from https://www.voidtools.com/forum/viewtopic.php?t=11081

**compile**

* `DOCKER_BUILDKIT=1 docker build --output . .`

**To use**

* `./make_efu /data "^/data" '\\server\data' > server.efu`
* on windows machine Everything Idexes > Filelist > choose \\server\server.efu
