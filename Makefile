
love:

fmt:
	gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*"`
	goimports -w `find . -type f -name '*.go' -not -path "./vendor/*"`

lint:
	golangci-lint run

ci:
	go test -race -coverprofile=coverage.txt -covermode=atomic
	[ -z "$CODECOV_TOKEN" ] || bash <(curl -s https://codecov.io/bash)
