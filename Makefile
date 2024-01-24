.PHONY: help run/docker test

all: build/docker test

APP = virtualbookstore

help: ## Print help for each target
	$(info Virtual Bookstore Makefile help.)
	$(info ====================================)
	$(info )
	$(info Available commands:)
	$(info )
	@grep '^[[:alnum:]_/]*:.* ##' $(MAKEFILE_LIST) \
		| sort | awk 'BEGIN {FS=":.* ## "}; {printf "%-25s %s\n", $$1, $$2};'

build/docker: ## Build project with docker
	 @DOCKER_BUILDKIT=1 docker build -t $(APP) .

run/docker: build/docker ## Run project with docker
	@docker run --rm --name $(APP) $(APP)

test: ## Run all tests
	@go test -v -failfast -cover ./internal/...

repl: ## Evaluate Golang expressions using REPL (Read-Eval-Print-Loop)
	@if type gore >/dev/null 2>&1 ; then gore; \
		else echo "SKIPPED. Run 'go install github.com/x-motemen/gore/cmd/gore@latest' and 'go install github.com/mdempsky/gocode@latest' first." >&2; \
	fi
