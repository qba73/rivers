.PHONY: help check cover test tidy

define PRINT_HELP_PYSCRIPT
import re, sys

for line in sys.stdin:
	match = re.match(r'^([a-zA-Z_-]+):.*?## (.*)$$', line)
	if match:
		target, help = match.groups()
		print("%-20s %s" % (target, help))
endef
export PRINT_HELP_PYSCRIPT

default: help

help:
	@python -c "$$PRINT_HELP_PYSCRIPT" < $(MAKEFILE_LIST)

check: ## Run static check analyzer
	staticcheck ./...

cover: ## Run unit tests and generate test coverage report
	go test -v ./... -count=1 -coverprofile=coverage.out
	go tool cover -html coverage.out
	staticcheck ./...

test: ## Run unit tests
	go test ./internal/... -v

# MODULES
tidy: ## Run go mod tidy and vendor
	go mod tidy
	go mod vendor
