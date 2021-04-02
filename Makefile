demo:
	go run cmd/okdoc/main.go --path "./examples/file1.star" --out ./examples/okdoc.json

test:
	go test -v -timeout 120s -race -bench=. -run=. ./...
