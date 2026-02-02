.PHONY: test cover bench check format clean

test:
	go test -cover ./...

cover:
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
	rm cover.out

bench:
	go test -bench=. ./...

check:
	go vet ./...

format:
	goimports -w .

clean:
	rm -f cover.html
