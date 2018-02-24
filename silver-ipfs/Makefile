PROJECT_NAME=silver-ipfs
CUR_DIR=$(CURDIR)
BIN_DIR=$(CUR_DIR)/bin

# Go parameters
GO_CMD=godep go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get

# Current time
CURRENT_TIME=`date "+%Y/%m/%d %H:%M:%S"`

all: build

build:
	$(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) -v $(CUR_DIR)/*.go
	@echo "$(CURRENT_TIME) [INFO ] Build completed\n"

clean:
	$(GO_CLEAN)
	rm $(BIN_DIR)/$(PROJECT_NAME)
	@echo "$(CURRENT_TIME) [INFO ] Clean completed\n"

test:
	@echo "Todo"

godep:
	godep save
	@echo "$(CURRENT_TIME) [INFO ] Godep saved\n"

run:
	@echo $(CURRENT_TIME) [INFO ] CUR_DIR=\"$(CUR_DIR)\"
	@echo $(CURRENT_TIME) [INFO ] BIN_DIR=\"$(BIN_DIR)\"
	$(BIN_DIR)/$(PROJECT_NAME)
