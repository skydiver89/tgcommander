VERSION := $(shell git describe --tags --abbrev=0)
GITREV := $(shell git describe --tags --dirty)
BUILDTIME := $(shell date)
LDFLAGS := '-w -s -X "main.VERSION=$(VERSION)" -X "main.GITREV=$(GITREV)" -X "main.BUILDTIME=$(BUILDTIME)"'
all:
	CGO_ENABLED=0 go build -ldflags=$(LDFLAGS)
