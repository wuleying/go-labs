# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
PROJECT_NAME=silver-blockchain
SRC_DIR=$(CURDIR)/src
BIN_DIR=$(CURDIR)/bin
DB_DIR=$(CURDIR)/db

all: build

build:
	$(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) -v $(SRC_DIR)/*.go
	@echo "Build completed"

run:
	@echo [INFO] SRC_DIR=$(SRC_DIR)
	@echo [INFO] BIN_DIR=$(BIN_DIR)
	$(BIN_DIR)/$(PROJECT_NAME) w c
	$(BIN_DIR)/$(PROJECT_NAME) w c

clean:
	$(GO_CLEAN)
	rm $(BIN_DIR)/$(PROJECT_NAME)
	rm $(DB_DIR)/$(PROJECT_NAME)-*.db
	rm $(DB_DIR)/wallet.dat
	@echo "Clean completed"