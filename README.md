### To run DB migration, from sql/schema run

- `goose postgres "postgres://postgres:postgres@localhost:5432/chirpy" up`

### To generate new models after creating/updating queries run

- `sqlc generate`

### To connect to chirpy local DB:

- `psql "postgres://postgres:postgres@localhost:5432/chirpy"`
