#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o ./build/
echo "linux版本编译完成"
GOOS=windows GOARCH=amd64 go build -o ./build/
echo "windows版本编译完成"
