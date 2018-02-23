PROJECT_NAME=silver-ipfs
CUR_DIR=$(CURDIR)
BIN_DIR=$(CUR_DIR)/bin

# Go parameters
GO_CMD=godep go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get

all: build

build:
	$(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) -v $(CUR_DIR)/*.go
	@echo "[INFO] Build completed"

clean:
	$(GO_CLEAN)
	rm $(BIN_DIR)/$(PROJECT_NAME)
	@echo "[INFO] Clean completed"

test:
	@echo "Todo"

godep:
	godep save
	@echo "[INFO] Godep saved"

run:
	@echo [INFO] CUR_DIR=\"$(CUR_DIR)\"
	@echo [INFO] BIN_DIR=\"$(BIN_DIR)\"
	$(BIN_DIR)/$(PROJECT_NAME)
