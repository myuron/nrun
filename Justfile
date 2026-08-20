default:
  @just --list

run:
  go run ./...

fmt:
  go fmt ./...

fmt-ci:
  nix fmt -- --ci

lint:
  golangci-lint run ./...

test:
  go test ./...

check: fmt lint
