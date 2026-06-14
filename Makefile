all: build

build:
	go build -o lesh ./cmd/lesh

clean:
	rm -f lesh

run:
	go run ./cmd/lesh

fmt:
	go fmt ./...

test:
	go test ./...
