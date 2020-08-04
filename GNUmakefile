TEST?="./provider"
PKG_NAME=wapc

default: build

build:
	go install

test:
	go test $(TEST) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(ACCTEST) -v -parallel 20 $(TESTARGS) -timeout 120m

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./$(PKG_NAME)

.PHONY: build test testacc fmt
