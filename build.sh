#!/usr/local/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=O go build -o churchws main.go
