PROJECT_NAME := gowatcher

VERSION := v0.0.1

ifeq (, $(shell which go))
$(error "No 'go' binary in $(PATH), consider rechecking your dev environment setup")
endif
GO_BIN   := go

ifeq (, $(shell which mockery))
$(error "No 'mockery' binary in $(PATH), consider installing https://github.com/vektra/mockery")
endif
MOCKERY_BIN   := mockery

PROJECT_DIR  := $(CURDIR)
BIN_NAME   := $(PROJECT_NAME)
GO_MODULE_NAME := git.int.kn/ibdev_go/$(PROJECT_NAME)
PROJECT_BUILD_DIR := $(PROJECT_DIR)/build

mainfilename := $(shell printf '%s' "${PROJECT_NAME}" | tr A-Z a-z)

.DEFAULT_GOAL := test

TEST_PACKAGE=
TEST_CASE=
GO_ARGS=
TEST_VERBOSE=
.PHONY: test
test:
	@$(GO_BIN) test $(if $(GO_ARGS),$(GO_ARGS) )$(if $(TEST_VERBOSE), -v )$(if $(TEST_PACKAGE), $(TEST_PACKAGE), $(PROJECT_DIR)/...) $(if $(TEST_CASE), -run "^$(TEST_CASE)$$", )


# go build -ldflags="-X 'main.Version=v1.0.0'"
RELEASE:=
INJECT_LOCATION:=github.com/saphieron/gowatcher/loop.Version
VERSION_INJECT:=-ldflags='-X "$(INJECT_LOCATION)=$(VERSION)"'
.PHONY: build
build: $(PROJECT_BUILD_DIR)
	@printf "Compiling '%s' to '$(PROJECT_BUILD_DIR)/$(BIN_NAME)'\n" $(PROJECT_DIR)/$(mainfilename).go
	@$(GO_BIN) build $(if $(RELEASE),$(VERSION_INJECT),) -o $(PROJECT_BUILD_DIR)/$(BIN_NAME) $(PROJECT_DIR)/$(mainfilename).go
	@printf "Done\n"

defaultRunArgs:=-v -n0.5 "ls -la"
RUN_ARGS=$(defaultRunArgs)
.PHONY: run
run: build
	@chmod +x $(PROJECT_BUILD_DIR)/$(BIN_NAME)
	@printf "Executing %s\n\n" "$(PROJECT_BUILD_DIR)/$(BIN_NAME)"
	@$(PROJECT_BUILD_DIR)/$(BIN_NAME) $(if $(RUN_ARGS),$(RUN_ARGS))


$(PROJECT_BUILD_DIR):
	@mkdir -m755 $@

.PHONY: cover
cover: $(PROJECT_BUILD_DIR)
	@$(GO_BIN) test -cover -covermode=count -coverprofile=$(PROJECT_BUILD_DIR)/profile.cov \
	$(if $(TEST_FILES), $(TEST_FILES), ./...)
	@$(GO_BIN) tool cover -func $(PROJECT_BUILD_DIR)/profile.cov

.PHONY: mock
mock:
	$(MOCKERY_BIN)