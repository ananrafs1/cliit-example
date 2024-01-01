.DEFAULT_GOAL := build

build :
	@go build -o ./example.exe main.go

.PHONY: build