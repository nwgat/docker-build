makemkv builder for ubuntu

* `docker compose build` (build)
* `docker compose up -d` (run)
* `id=$(docker create makemkv:1.18.2) && docker cp $id:/tmp/ . && docker rm -v $id` (copy debs)
