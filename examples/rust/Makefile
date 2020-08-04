COLOR ?= always # Valid COLOR options: {always, auto, never}
CARGO = cargo --color $(COLOR)
DEBUG = $(TARGET)/debug
RELEASE = $(TARGET)/release

.PHONY: all build check clean test update

all: build

build:
	@$(CARGO) build --target wasm32-unknown-unknown

check:
	@$(CARGO) check

clean:
	@$(CARGO) clean

test: build
	@$(CARGO) test

update:
	@$(CARGO) update

release:
	@$(CARGO) build --release --target wasm32-unknown-unknown
