GIT_REPO=github.com/amukherj/samplegorpc
rdir := $(dir $(lastword $(MAKEFILE_LIST)))
BINARIES=$(rdir)/bin/gorpcsvr $(rdir)/bin/gorpcclient $(rdir)/bin/jsonrpc

.PHONY: test clean binaries

all: binaries

binaries: $(BINARIES)

$(rdir)/bin/gorpcsvr:
	mkdir -p $(rdir)/bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ $(GIT_REPO)/cmd/gorpcsvr

$(rdir)/bin/gorpcclient:
	mkdir -p $(rdir)/bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ $(GIT_REPO)/cmd/gorpcclient

$(rdir)/bin/jsonrpc:
	mkdir -p $(rdir)/bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ $(GIT_REPO)/cmd/jsonrpc

clean:
	rm -f $(BINARIES)
