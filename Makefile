# Connect to the primary database
.PHONY: db
db:
	docker compose exec -it db psql postgresql://postgres:password@localhost:5432/critter

# Connect to the test database
.PHONY: db-test
db-test:
	docker compose exec -it db psql postgresql://postgres:password@localhost:5432/critter_test

# Connect to the primary cache
.PHONY: cache
cache:
	docker compose exec -it cache redis-cli

 # Connect to the test cache
.PHONY: cache-test
cache-test:
	docker compose exec -it redis-cli -n 1

# Generate Ent code
.PHONY: generate
ent-gen:
	go generate ./ent

# Start the Docker containers & remove old volumes
.PHONY: up
up:
	docker compose up -d

# Stop the Docker containers
.PHONY: down
down:
	docker compose down

# Rebuild Docker containers to wipe all data
.PHONY: reset
reset: down
	make up

# Run the application
.PHONY: run
run: up
	clear
	go run main.go

# Run all tests
.PHONY: test
test:
	go test -count=1 -p 1 ./...

# Run the indexer
.PHONY: indexer
indexer:
	clear
	go run indexer/indexer.go

# Run the worker
.PHONY: worker
worker:
	clear
	go run worker/worker.go

# Check for direct dependency updates
.PHONY: update
check-updates:
	go list -u -m -f '{{if not .Indirect}}{{.}}{{end}}' all | rg "\["
