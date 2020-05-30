## letsgopokemon

# start local dev

In the terminal:

- Start the docker container `docker-compose up -d`

If you are working for the first time or have removed your docker volumes

- Create the database `docker exec -it letsgopokemon_pg_1 bash`, `psql -U letsgo`,
  `CREATE DATABASE letsgopokemon`, `\q`, `exit`

Source the env variables

- `source ./env`

# generate graphql based on schema

If changes are made to the GraphQL schema:

`go run github.com/99designs/gqlgen -v`

# updating folder structure

If changes are made to the folder structure, update `gqlgen.yml` to ensure schema generation is still pointing to the correct files.

# to run locally

To start GraphQL playground locally:

`go run ./server.go`

It can be accessed on `http://localhost:8080`
