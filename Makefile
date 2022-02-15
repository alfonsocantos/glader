PROJECT = github.com/alfonsocantos/glader

GOCMD=go

.PHONY: deps test

all: test

deps:
	$(GOCMD) mod tidy -v
	$(GOCMD) mod download
	$(GOCMD) mod vendor

test:
	$(GOCMD) test $(PROJECT)/... -cover
