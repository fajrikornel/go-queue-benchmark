
build: clean test
	go build -o build/ ./cmd/main.go

test:
	go test ./... -v

clean:
	rm -rf ./build

run: build
	./build/main
