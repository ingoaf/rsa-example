all: clean imports lint vet test
.PHONY: all

imports:
	go get -u golang.org/x/tools/cmd/goimports
	goimports -w .

lint:
	go get -u golang.org/x/lint/golint
	golint ./...

vet:
	go vet ./...

test:
	go test ./...

build:
	go mod download
	CGO_ENABLED=0 GOARCH=amd64 GO111MODULE=on \
	go build -a -installsuffix cgo -o app main.go

run:
	go run main.go

clean:
	go clean -modcache
