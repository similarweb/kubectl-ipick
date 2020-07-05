# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD)fmt
GOARCH=$(shell go env GOARCH)

BINARY_NAME=kubectl-interactive

TEST_EXEC_CMD=$(GOTEST) -coverprofile=cover.out -short -cover -failfast ./... 

test: ## Run tests for the project
		$(TEST_EXEC_CMD)

lint: ## Code linting
	golangci-lint run

fmt: ## Validate go format
	@echo checking gofmt...
	@res=$$($(GOFMT) -d -e -s $$(find . -type d \( -path ./src/vendor \) -prune -o -name '*.go' -print)); \
	if [ -n "$${res}" ]; then \
		echo checking gofmt fail... ; \
		echo "$${res}"; \
		exit 1; \
	else \
		echo Your code formating is according gofmt standards; \
	fi
	
release:  releasebin ## Build and release all platforms builds to nexus

releasebin: ## Create release with platforms
	goreleaser build
		
build-linux: ## Build Cross Platform Binary
		CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) $(GOBUILD) -o $(BINARY_NAME)_linux -v

build-osx: ## Build Mac Binary
		CGO_ENABLED=0 GOOS=darwin GOARCH=$(GOARCH) $(GOBUILD) -o $(BINARY_NAME)_osx -v

build-windows: ## Build Windows Binary
		CGO_ENABLED=0 GOOS=windows GOARCH=$(GOARCH) $(GOBUILD) -o $(BINARY_NAME)_windows -v

build-docker: ## BUild Docker image file
		$(DOCKER) build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

help: ## Show Help menu
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
