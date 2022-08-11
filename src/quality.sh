#!/bin/sh
echo "starting ..."
echo "adding needed deps"
go get github.com/golang/mock/mockgen@v1.5.0
echo "running go generate"
go generate ./...
echo "running go fmt"
go fmt ./...
echo "do a go mod tidy"
go mod tidy
echo "running golangci-lint"
golangci-lint run ./... -v
echo "running tests with coverage flag"
# shellcheck disable=SC2006
go test `go list ./... | grep -v example` -coverprofile=coverage.txt -covermode=atomic
echo "I don finish, all izz well?"
echo "abeg confirm !!! No Error?"
