REPO_ROOT	:= $(shell git rev-parse --show-toplevel)

.DEFAULT_GOAL:=help
SHELL:=/usr/bin/env bash

COLOR:=\\033[36m
NOCOLOR:=\\033[0m
GITREPO=$(shell git remote -v | grep fetch | awk '{print $$2}' | sed 's/\.git//g' | sed 's/https:\/\///g')

PROJECTS=  nft-meta block-etl image-converter
GO_PROJECTS=  nft-meta block-etl

##@ init project
init:
	cp -f .githooks/* .git/hooks

go.mod:
	go mod init ${GITREPO}
	go mod tidy -compat=1.17

deps:
	go get -d ./...
	go mod tidy -compat=1.17

##@ Verify

.PHONY: add-verify-hook verify verify-build verify-golangci-lint verify-go-mod verify-shellcheck verify-spelling all

add-verify-hook: ## Adds verify scripts to git pre-commit hooks.
# Note: The pre-commit hooks can be bypassed by using the flag --no-verify when
# performing a git commit.
	git config --local core.hooksPath "${REPO_ROOT}/.githooks"

# TODO(lint): Uncomment verify-shellcheck once we finish shellchecking the repo.
verify: go.mod verify-golangci-lint verify-go-mod #verify-shellcheck ## Runs verification scripts to ensure correct execution
	${REPO_ROOT}/hack/verify.sh

verify-build: ## Build project
	@for x in $(PROJECTS); do \
		${REPO_ROOT}/$${x}/script/build.sh $${x};\
	done

verify-go-mod: ## Runs the go module linter
	${REPO_ROOT}/hack/verify-go-mod.sh

verify-golangci-lint: ## Runs all golang linters
	${REPO_ROOT}/hack/verify-golangci-lint.sh

verify-shellcheck: ## Runs shellcheck
	${REPO_ROOT}/hack/verify-shellcheck.sh

verify-spelling: ## Verifies spelling.
	${REPO_ROOT}/hack/verify-spelling.sh


gen-ent:
	go get entgo.io/ent/cmd/ent@v0.11.2
	go run entgo.io/ent/cmd/ent generate --feature entql,sql/upsert,privacy,schema/snapshot,sql/modifier ./nft-meta/pkg/db/ent/schema

ifdef AIMPROJECT
PROJECTS= $(AIMPROJECT)
endif

ifndef DEVELOPMENT
DEVELOPMENT= dev
endif

ifndef DOCKER_REGISTRY
DOCKER_REGISTRY= x
endif

ifndef TAG
TAG= latest
endif
	
.PHONY: build-docker release-docker deploy-to-k8s-cluster

build-docker:
	@for x in $(PROJECTS); do \
		${REPO_ROOT}/$${x}/script/build-docker-image.sh $(DEVELOPMENT) $(DOCKER_REGISTRY);\
	done
release-docker:
	@for x in $(PROJECTS); do \
		${REPO_ROOT}/$${x}/script/release-docker-image.sh $(TAG) $(DOCKER_REGISTRY);\
	done
deploy-to-k8s-cluster:
	@for x in $(PROJECTS); do \
		${REPO_ROOT}/$${x}/script/deploy-to-k8s-cluster.sh $(TAG);\
	done
	 

##@ Tests

.PHONY: go-unit-test go-ut
go-ut: unit-test
go-unit-test: verify-build
	@for x in $(GO_PROJECTS); do \
		${REPO_ROOT}/$${x}/script/before-test.sh;\
	done
	@for x in $(GO_PROJECTS); do \
		${REPO_ROOT}/$${x}/script/test-go.sh;\
	done
	@for x in $(GO_PROJECTS); do \
		${REPO_ROOT}/$${x}/script/after-test.sh;\
	done

test-verbose:
	VERBOSE=1 make test

##@ Helpers

.PHONY: help

help:  ## Display this help
	@awk \
		-v "col=${COLOR}" -v "nocol=${NOCOLOR}" \
		' \
			BEGIN { \
				FS = ":.*##" ; \
				printf "\nUsage:\n  make %s<target>%s\n", col, nocol \
			} \
			/^[a-zA-Z_-]+:.*?##/ { \
				printf "  %s%-15s%s %s\n", col, $$1, nocol, $$2 \
			} \
			/^##@/ { \
				printf "\n%s%s%s\n", col, substr($$0, 5), nocol \
			} \
		' $(MAKEFILE_LIST)
