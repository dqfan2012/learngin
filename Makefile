# Makefile

.PHONY: all server migrate-up migrate-down make-migration model

# Default target
all: server

# Run the web server
server:
	go run ./cmd/server/main.go

# Apply all up migrations
migrate-up:
	go run ./cmd/migrate/main.go -up

# Rollback all migrations
migrate-down:
	go run ./cmd/migrate/main.go -down

# Create new migration files
migration:
	go run ./cmd/migrate/main.go -new $(name)

# Create a new model
# I.E., make model name=Blog/Post
model:
	go run ./cmd/model/main.go -name=$(name)

# Seed the database
seed:
	go run ./cmd/seeder/main.go

# Automatically rollback migrations, rerun migrations, seed the db, and run the server
setup-dev: migrate-down migrate-up seed server
