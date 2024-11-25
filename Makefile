.PHONY: all build tidy clean

targets := $(wildcard mwtest/*.go) main.go

default: all

all: tidy build

tidy: go.mod
	@echo "Tidying.."
	@go mod tidy

build: mwt

clean:
	@echo "Cleaning..."
	@rm -f mwt
	@go clean

mwt: $(targets) go.mod go.sum
	@echo "Building..."
	@go build -o mwt main.go

