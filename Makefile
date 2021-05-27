
love:

fmt:
	gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*"`
	goimports -w `find . -type f -name '*.go' -not -path "./vendor/*"`

lint:
	golangci-lint run

ci: cover
	[ -z "$(CODECOV_TOKEN)" ] || bash -c 'bash <(curl -s https://codecov.io/bash)'

.PHONY: cover
cover:
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html
