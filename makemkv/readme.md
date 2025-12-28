makemkv builder for ubuntu

* `docker compose build` (build)
* `docker compose up -d` (run)
* `docker exec -it makemkv makemkvcon` (run inside the docker container)
* `docker cp makemkv:/tmp/ .` (copy the debs to install locally)
* `id=$(docker create makemkv:1.18.2) && docker cp $id:/tmp/ . && docker rm -v $id` (alternative way to copy debs)
