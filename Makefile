.PHONY: build
build:
	@ go build -o ./bin/glox ./cmd/glox

.PHONY: test
test:
	@ go test ./...
