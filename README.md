## letsgopokemon

# start local dev

To start, run `docker-compose up -d` in the terminal to create the database

# generate graphql based on schema

If changes are made to the GraphQL schema:

`go run github.com/99designs/gqlgen -v`

# updating folder structure

If changes are made to the folder structure, update `gqlgen.yml` to ensure schema generation is still pointing to the correct files.
