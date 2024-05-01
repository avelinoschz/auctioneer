GENERATED_TARGET_DIR ?= $(CURDIR)/generated
BUILD_TARGET_DIR ?= $(GENERATED_TARGET_DIR)/bin

GO_CODE_COVERAGE_OUTPUT := code-coverage.out
GO_TEST_FLAGS := -v -shuffle=on -count=1 -timeout=10s -coverprofile=$(GO_CODE_COVERAGE_OUTPUT)
GO_BUILD_FLAGS := -ldflags="-w -s"

DOCKER_IMAGE_TAG := auctioneer

.PHONY: build-go
build-go:
	CGO_ENABLED=0 go build $(GO_BUILD_FLAGS) -o $(BUILD_TARGET_DIR)/auctioneer $(CURDIR)/cmd/

.PHONY: run-go
run-go: build-go
	$(BUILD_TARGET_DIR)/auctioneer

.PHONY: run-test
run-test:
	go test $(CURDIR)/pkg/auctioneer $(GO_TEST_FLAGS)

.PHONY: open-cover
open-cover: run-test
	go tool cover -html=$(GO_CODE_COVERAGE_OUTPUT)

.PHONY: ensure-docker
ensure-docker:
	@if ! docker info >/dev/null 2>&1; then \
		echo "Docker is not available. Please ensure Docker is installed and running."; \
		exit 1; \
	fi

.PHONY: build-docker
build-docker: ensure-docker
	docker build -t $(DOCKER_IMAGE_TAG) .

.PHONY: run-docker
run-docker: ensure-docker
	docker run --rm $(DOCKER_IMAGE_TAG) 