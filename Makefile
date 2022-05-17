.PHONY: install test build serve clean pack deploy ship
TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get .

test: install
	go test ./...

build: install
	go build -ldflags "-X main.version=$(TAG)" -o ./bin/run-platform-service .

serve: build
	./bin/run-platform-service serve grpc

clean:
	rm -f ./bin/run-platform-service

dev:
	GOOS=linux make build
	docker build -t localhost:5000/run-platform-service .
	docker push localhost:5000/run-platform-service