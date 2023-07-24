PROJECT_NAME := gowatcher

ifeq (, $(shell which go))
$(error "No 'go' binary in $(PATH), consider rechecking your dev environment setup")
endif
GO_BIN   := go

PROJECT_DIR  := $(CURDIR)
BIN_NAME   := $(PROJECT_NAME)
GO_MODULE_NAME := git.int.kn/ibdev_go/$(PROJECT_NAME)


mainfilename := $(shell printf '%s' "${PROJECT_NAME}" | tr A-Z a-z)

PROJECT_BUILD_DIR := $(PROJECT_DIR)/build

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