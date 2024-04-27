# Before using, configure goproxy; see https://goproxy.githubapp.com/setup.
.PHONY: all
all: build

.PHONY: build
build: tools twirp

tools:
	go install github.com/golang/protobuf/protoc-gen-go@v1.5.2
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2
	.devcontainer/scripts/install-tools.sh

.PHONY: test
test:  test-go

.PHONY: tools.ruby
tools.ruby:
	.devcontainer/scripts/install-ruby-tools.sh

.PHONY: lint
lint: tools vet
	@echo "==> linting Go code <=="
	@golangci-lint run ./...

.PHONY: vet
vet:
	@echo "==> vetting Go code <=="
	go vet ./...

.PHONY: proto
proto:
	script/helpers/protoc

.PHONY: clean
clean:
	rm -f tools
	go clean -cache -testcache -modcache

.PHONY: twirp
twirp:
	go build -o bin/twirp-service ./cmd/server

test-go:
	@echo "==> running Go tests <=="
	CGO_ENABLED=1 go test -p 64 -race ./...
