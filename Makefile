
run: build
	./build/main

build: clean
	go build -o build/ ./cmd/main.go

clean:
	rm -rf ./build
