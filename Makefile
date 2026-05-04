GO:=$(shell which go)

all:
	$(GO) run cmd/api/main.go