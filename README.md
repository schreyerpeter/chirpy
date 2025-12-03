## About Chirpy

- Chirpy is a simple clone of Twitter's basic functionality, built with Go
- Basic auth and resource management endpoints are included

## Development Prompts

### To run DB migration, from sql/schema:

- `goose postgres "postgres://postgres:postgres@localhost:5432/chirpy" up`

### To generate new models after creating/updating queries:

- `sqlc generate`

### To connect to chirpy local DB:

- `psql "postgres://postgres:postgres@localhost:5432/chirpy"`
