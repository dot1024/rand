
build:
	@go build -o build/rand-gen rand-gen/main.go

.PHONY: all build bin clean
