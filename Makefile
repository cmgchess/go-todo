APP_NAME=gotodo
ENTRY=cmd/api/main.go
BIN_DIR=bin

run:
	go run $(ENTRY)

build: $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(ENTRY)

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

clean:
	rm -rf $(BIN_DIR)

.PHONY: run build clean
