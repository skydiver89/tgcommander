VERSION := $(shell git describe --tags --abbrev=0)
GITREV := $(shell git describe --tags --dirty)
BUILDTIME := $(shell date +"%H:%M:%S %d/%m/%Y")
LDFLAGS := '-w -s -X "main.VERSION=$(VERSION)" -X "main.GITREV=$(GITREV)" -X "main.BUILDTIME=$(BUILDTIME)"'
.PHONY: all deb
.ONESHELL:
all:
	mkdir -p build
	CGO_ENABLED=0 go build -ldflags=$(LDFLAGS) -o ./build/tgcommander_linux_x86_64
deb: all
	./buildscripts/builddeb.sh
