PROJECT_NAME=silver-ipfs
SRC_DIR=$(CURDIR)/src
BIN_DIR=$(CURDIR)/bin

# Go parameters
GO_CMD=go
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
	@echo "Clean completed"

test:
	@echo "Todo"

run:
	@echo [INFO] SRC_DIR=$(SRC_DIR)
	@echo [INFO] BIN_DIR=$(BIN_DIR)
	$(BIN_DIR)/$(PROJECT_NAME)
