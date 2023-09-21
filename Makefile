PROJECT_NAME := gowatcher

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



.DEFAULT_GOAL := build

TEST_PACKAGE=
TEST_CASE=
GO_ARGS=
TEST_VERBOSE=
.PHONY: test
test:
	@$(GO_BIN) test $(if $(GO_ARGS),$(GO_ARGS) )$(if $(TEST_VERBOSE), -v )$(if $(TEST_PACKAGE), $(TEST_PACKAGE), $(PROJECT_DIR)/...) $(if $(TEST_CASE), -run "^$(TEST_CASE)$$", )


.PHONY: build
build: $(PROJECT_BUILD_DIR)
	@printf "Compiling '%s' to '$(PROJECT_BUILD_DIR)/$(BIN_NAME)'\n" $(PROJECT_DIR)/$(mainfilename).go
	@$(GO_BIN) build -o $(PROJECT_BUILD_DIR)/$(BIN_NAME) $(PROJECT_DIR)/$(mainfilename).go
	@printf "Done\n"

RUN_ARGS=
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