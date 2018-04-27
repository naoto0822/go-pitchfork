TARGET := ./pitchfork

## all task
all: dep vet lint test

## install dependencies
dep:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/lint/golint
	go get github.com/Songmu/make2help/cmd/make2help

## run vet
vet:
	go vet $(TARGET)

## run lint
lint:
	golint -set_exit_status $(TARGET)

## run test
test:
	go test -v $(TARGET)

## show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: dep vet lint test help