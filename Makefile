APP_NAME=gotodo
ENTRY=cmd/api/main.go
BIN_DIR=bin

run:
	@go run $(ENTRY)

build: $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(ENTRY)

$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

clean:
	@rm -rf $(BIN_DIR)

test:
	@go test -v -cover ./...

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

.PHONY: run build clean test migration migrate-up migrate-down