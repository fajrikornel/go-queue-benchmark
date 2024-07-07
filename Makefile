
bench:
	go test -bench=. ./cmd

test:
	go test ./... -v
