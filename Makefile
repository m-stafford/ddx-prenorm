.DEFAULT_GOAL := build

.PHONY:fmt vet build
clean:
				rm prenorm

fmt:
				go fmt ./...

vet: fmt
				go vet ./...

build: vet
				go build
