.PHONY: all fmt build clean run check cover lint docker help
GOFMT_FILES?=$$(find . -name '*.go')

all: fmt build

fmt:
	gofmt -w $(GOFMT_FILES)

lint:
	go vet

image:
	docker build -t webhookbd .
