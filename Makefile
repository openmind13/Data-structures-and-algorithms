start:
	go build -o main *.go && ./main

test:
	go test -v ./...

.DEFAULT_GOAL := start