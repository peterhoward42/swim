
GOLANGCI_LINT_DEP=github.com/golangci/golangci-lint/cmd/golangci-lint@v1.16.0

.PHONY: test
test: style
test:
	mkdir -p ./render/testresults/new
	go test ./...

.PHONY: deps
deps: ## Install dependencies
	go get
	go get $(GOLANGCI_LINT_DEP)

.PHONY: fmt
fmt: build
	gofmt -w -s -l .

.PHONY: style
style: fmt
	golangci-lint run --enable=gofmt -v

.PHONY: build
build: deps
	go build  ./...


