build:
	go build -v -o ./bin/rotator ./cmd/rotator

test:
	go test -race -count=100 ./...

install-lint:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh

lint: install-lint
	golangci-lint run ./...

run:
	docker-compose -f ./deployment/docker-compose.yaml -p rotator up --build

stop:
	docker-compose -f ./deployment/docker-compose.yaml -p rotator down

.PHONY: build test install-lint lint run stop