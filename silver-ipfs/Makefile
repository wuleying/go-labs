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
CUR_TIME=`date "+%Y/%m/%d %H:%M:%S"`

all: build

build:
	$(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) -v $(CUR_DIR)/*.go
	@echo "$(CUR_TIME) [INFO ] Build completed"

clean:
	$(GO_CLEAN)
	rm $(BIN_DIR)/$(PROJECT_NAME)
	@echo "$(CUR_TIME) [INFO ] Clean completed"

test:
	$(GO_TEST)

test_cover:
	$(GO_TEST) -cover

godep:
	godep save
	@echo "$(CUR_TIME) [INFO ] Godep saved"

ps:
	ps -ef | grep $(PROJECT_NAME)

run:
	@echo $(CUR_TIME) [INFO ] CUR_DIR=\"$(CUR_DIR)\"
	@echo $(CUR_TIME) [INFO ] BIN_DIR=\"$(BIN_DIR)\"
	$(BIN_DIR)/$(PROJECT_NAME)

stop:
	pgrep -f $(PROJECT_NAME) | xargs kill -9