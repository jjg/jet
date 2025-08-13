#!/bin/bash
# build jet
GOOS=linux GOARCH=amd64 go build -o bin/jet-linux-amd64 ./tools/jet/
GOOS=linux GOARCH=arm64 go build -o bin/jet-linux-arm64 ./tools/jet/
GOOS=linux GOARCH=arm go build -o bin/jet-linux-arm ./tools/jet/
GOOS=darwin GOARCH=amd64 go build -o bin/jet-darwin-amd64 ./tools/jet/
GOOS=darwin GOARCH=arm64 go build -o bin/jet-darwin-arm64 ./tools/jet/
GOOS=windows GOARCH=amd64 go build -o bin/jet-windows-amd64 ./tools/jet/
# build today
GOOS=linux GOARCH=amd64 go build -o bin/today-linux-amd64 ./tools/today/
GOOS=linux GOARCH=arm64 go build -o bin/today-linux-arm64 ./tools/today/
GOOS=linux GOARCH=arm go build -o bin/today-linux-arm ./tools/today/
GOOS=darwin GOARCH=amd64 go build -o bin/today-darwin-amd64 ./tools/today/
GOOS=darwin GOARCH=arm64 go build -o bin/today-darwin-arm64 ./tools/today/
GOOS=windows GOARCH=amd64 go build -o bin/today-windows-amd64 ./tools/today/
# build yesterday
GOOS=linux GOARCH=amd64 go build -o bin/yesterday-linux-amd64 ./tools/yesterday/
GOOS=linux GOARCH=arm64 go build -o bin/yesterday-linux-arm64 ./tools/yesterday/
GOOS=linux GOARCH=arm go build -o bin/yesterday-linux-arm ./tools/yesterday/
GOOS=darwin GOARCH=amd64 go build -o bin/yesterday-darwin-amd64 ./tools/yesterday/
GOOS=darwin GOARCH=arm64 go build -o bin/yesterday-darwin-arm64 ./tools/yesterday/
GOOS=windows GOARCH=amd64 go build -o bin/yesterday-windows-amd64 ./tools/yesterday/
