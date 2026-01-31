.PHONY: test bench check format

test:
	go test ./...

bench:
	go test -bench=. ./...

check:
	go vet ./...

format:
	goimports -w .
