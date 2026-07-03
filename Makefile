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

tidy:
	grep -v '^.tea.kareha.org' go.mod > go.mod.clipped
	mv go.mod.clipped go.mod
	GOPRIVATE=tea.kareha.org go mod tidy
