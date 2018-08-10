PROJECT_NAME=silver-ui

CUR_DIR=$(CURDIR)
BIN_DIR=$(CUR_DIR)/bin

BRANCH=`git rev-parse --abbrev-ref HEAD`
SHA1=`git rev-parse --short HEAD`
CUR_DATE=`date +"%Y%m%d"`
CUR_TIME=`date "+%Y/%m/%d %H:%M:%S"`

VERSION=$(BRANCH).$(CUR_DATE).$(SHA1)

LDFLAGS=-ldflags "-X \"github.com/wuleying/go-labs/silver-ui/version.Version=$(VERSION)\""

# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_DOC=$(GO_CMD) doc
GO_GET=$(GO_CMD) get
GO_FMT=$(GO_CMD) fmt
GO_IMPORTS=goimports

# Tools
default: fmt build

dev: stop default

build:
	$(GO_BUILD) $(LDFLAGS) -o $(BIN_DIR)/$(PROJECT_NAME) -v $(CUR_DIR)/*.go
	@echo "$(CUR_TIME) [INFO ] Build $(PROJECT_NAME) completed"

clean:
	$(GO_CLEAN)
	rm $(BIN_DIR)/$(PROJECT_NAME)
	@echo "$(CUR_TIME) [INFO ] Clean completed"

fmt:
	$(GO_FMT) .
	@echo "$(CUR_TIME) [INFO ] Go fmt completed"

imports:
	$(GO_IMPORTS) $(shell find . -name "*.go" | egrep -v "vendor")
	@echo "$(CUR_TIME) [INFO ] Go imports completed"

ps:
	ps -ef | grep $(PROJECT_NAME)

run:
	@echo $(CUR_TIME) [INFO ] CUR_DIR=$(CUR_DIR)
	@echo $(CUR_TIME) [INFO ] BIN_DIR=$(BIN_DIR)
	@echo $(CUR_TIME) [INFO ] VERSION=$(VERSION)
	$(BIN_DIR)/$(PROJECT_NAME)

stop:
	pgrep -f $(PROJECT_NAME) | xargs kill -9

# Test tools
test:
	$(GO_TEST) ./utils

cover:
	$(GO_TEST) -cover ./utils

bench:
	$(GO_TEST) -bench=. ./utils

# Go docs
doc:
	$(GO_DOC) ./utils

# Check tools
check: vet lint gocyclo gosimple unused staticcheck ineffassign misspell

vet:
	@echo go vet
	@if test -n '$(shell go vet `glide nv` 2>&1)'; then \
		echo '$(shell go vet `glide nv` 2>&1)'; \
	fi

lint:
	@echo golint
	@if test -n '$(shell golint `glide nv` 2>&1)'; then \
		echo '$(shell golint `glide nv` 2>&1)'; \
	fi

gocyclo:
	gocyclo -over 20 $(shell find . -name "*.go" | egrep -v "vendor")

gosimple:
	gosimple $(shell glide nv)

unused:
	unused $(shell glide nv)

staticcheck:
	staticcheck $(shell glide nv)

ineffassign:
	@for f in `find . -type d -depth 1 | egrep -v "git|hook|vendor"`; do \
		ineffassign $$f; \
	done

misspell:
	misspell -i "unknwon" $(shell find . -maxdepth 1 -mindepth 1 -type d | egrep -v "vendor|doc|bin|.git|.idea")

goconst:
	goconst $(shell glide nv)

# Get tools and third packages
get:
	$(GO_GET) github.com/andlabs/ui

glide:
	glide install
	glide update