help:
	@echo "Please use 'make <target>' where <target> is one of the following:"
	@echo "  lint            to run the linter."
	@echo "  test            to run the tests."

.PHONY: tidy
tidy:
	go mod tidy

lint:
	golangci-lint run --modules-download-mode=vendor --timeout=2m0s -E golint --exclude-use-default=false --build-tags integration

test:
	GO111MODULE=on go test -mod=vendor `go list -mod vendor ./...`  -race

.PHONY: cover
cover:
	go test -race -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out

fmt:
	go fmt ./...

.PHONY: test