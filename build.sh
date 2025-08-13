#!/bin/bash
# build jet
GOOS=linux GOARCH=amd64 go build -o bin/jet-linux-amd64 ./cmd/jet/
GOOS=linux GOARCH=arm64 go build -o bin/jet-linux-arm64 ./cmd/jet/
GOOS=linux GOARCH=arm go build -o bin/jet-linux-arm ./cmd/jet/
GOOS=darwin GOARCH=amd64 go build -o bin/jet-darwin-amd64 ./cmd/jet/
GOOS=darwin GOARCH=arm64 go build -o bin/jet-darwin-arm64 ./cmd/jet/
GOOS=windows GOARCH=amd64 go build -o bin/jet-windows-amd64 ./cmd/jet/
# build today
GOOS=linux GOARCH=amd64 go build -o bin/today-linux-amd64 ./cmd/today/
GOOS=linux GOARCH=arm64 go build -o bin/today-linux-arm64 ./cmd/today/
GOOS=linux GOARCH=arm go build -o bin/today-linux-arm ./cmd/today/
GOOS=darwin GOARCH=amd64 go build -o bin/today-darwin-amd64 ./cmd/today/
GOOS=darwin GOARCH=arm64 go build -o bin/today-darwin-arm64 ./cmd/today/
GOOS=windows GOARCH=amd64 go build -o bin/today-windows-amd64 ./cmd/today/
# build yesterday
GOOS=linux GOARCH=amd64 go build -o bin/yesterday-linux-amd64 ./cmd/yesterday/
GOOS=linux GOARCH=arm64 go build -o bin/yesterday-linux-arm64 ./cmd/yesterday/
GOOS=linux GOARCH=arm go build -o bin/yesterday-linux-arm ./cmd/yesterday/
GOOS=darwin GOARCH=amd64 go build -o bin/yesterday-darwin-amd64 ./cmd/yesterday/
GOOS=darwin GOARCH=arm64 go build -o bin/yesterday-darwin-arm64 ./cmd/yesterday/
GOOS=windows GOARCH=amd64 go build -o bin/yesterday-windows-amd64 ./cmd/yesterday/
