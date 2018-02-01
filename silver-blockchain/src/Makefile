PROJECT_NAME=silver-blockchain
PACKAGE=go-labs/silver-blockchain
BASE_DIR=$(GOPATH)/src/$(PACKAGE)
SRC_DIR=$(BASE_DIR)/src
BIN_DIR=$(BASE_DIR)/bin
DB_DIR=$(BASE_DIR)/db

# Go parameters
GO_CMD=godep go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get

all: build

build:
	$(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) -v $(SRC_DIR)/*.go
	@echo "Build completed"

clean:
	$(GO_CLEAN)
	rm $(BIN_DIR)/$(PROJECT_NAME)
	rm $(DB_DIR)/wallet.dat
	rm $(DB_DIR)/$(PROJECT_NAME)-*.db
	@echo "Clean completed"

test:
	@echo "Todo"

run:
	@echo [INFO] SRC_DIR=$(SRC_DIR)
	@echo [INFO] BIN_DIR=$(BIN_DIR)
	$(BIN_DIR)/$(PROJECT_NAME)

wallte:
	$(BIN_DIR)/$(PROJECT_NAME) w c
