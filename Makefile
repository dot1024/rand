
build:
	go build -o build/rand-gen rand-gen/main.go

test:
	go test -bench=. -benchmem

.PHONY: all build bin clean
