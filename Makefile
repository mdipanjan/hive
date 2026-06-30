VERSION := 0.1.0
BUILD_FLAGS := -ldflags "-s -w -X main.Version=$(VERSION)"
CGO_FLAGS := CGO_CFLAGS="-w"

PRETTIER := npx --yes prettier@3
MARKDOWNLINT := npx --yes markdownlint-cli2@0.18

.PHONY: build install clean run fmt fmt-go fmt-md lint lint-go lint-md test

build:
	$(CGO_FLAGS) go build $(BUILD_FLAGS) -o hive .

## fmt: auto-format Go + Markdown
fmt: fmt-go fmt-md

fmt-go:
	gofmt -w .

fmt-md:
	$(PRETTIER) --write "**/*.md"
	$(MARKDOWNLINT) --fix "**/*.md" || true

## lint: check formatting + vet + markdown (no writes); used in CI
lint: lint-go lint-md

lint-go:
	@unformatted=$$(gofmt -l .); \
	if [ -n "$$unformatted" ]; then \
		echo "gofmt needs to run on:"; echo "$$unformatted"; exit 1; \
	fi
	$(CGO_FLAGS) go vet ./...

lint-md:
	$(PRETTIER) --check "**/*.md"
	$(MARKDOWNLINT) "**/*.md"

## test: run the test suite
test:
	$(CGO_FLAGS) go test ./...

install:
	$(CGO_FLAGS) go install $(BUILD_FLAGS) .

run: build
	./hive

clean:
	rm -f hive

version:
	@echo $(VERSION)
