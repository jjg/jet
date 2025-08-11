#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o bin/jet-linux-amd64
GOOS=linux GOARCH=arm64 go build -o bin/jet-linux-arm64
GOOS=linux GOARCH=arm go build -o bin/jet-linux-arm
GOOS=darwin GOARCH=amd64 go build -o bin/jet-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o bin/jet-darwin-arm64
GOOS=windows GOARCH=amd64 go build -o bin/jet-windows-amd64
