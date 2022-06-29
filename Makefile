fmt:
	go fmt

lint:
	go vet

build: fmt lint
	go build
