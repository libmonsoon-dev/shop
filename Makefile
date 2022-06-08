GO_BIN=go
PKG_PATTERN=./...
GO_FILES=$(shell GO_BIN=$(GO_BIN) PKG_PATTERN=$(PKG_PATTERN) ./get-go-file-names.sh)

pre-commit: check-hook lint-fix build test

watcher-hook: build test

check-hook: #TODO:
	@ls .git/hooks/pre-commit 2> /dev/null || \
	 (echo "#!/bin/sh\nmake pre-commit" > .git/hooks/pre-commit \
	 && chmod +x .git/hooks/pre-commit)

lint-fix: goimports

goimports:
	@for f in $(GO_FILES) ; do \
		echo goimports $$f; \
		goimports -w -force-rewrite -local github.com/libmonsoon-dev $$f ; \
	done

build:
	$(GO_BIN) build -v $(PKG_PATTERN)

test:
	$(GO_BIN) test -v $(PKG_PATTERN)
