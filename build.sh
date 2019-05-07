#!/bin/bash

PROJECT_NAME="kafkaprod"

echo "build started"

mkdir build

echo "building linux-x64 version"
env GOOS=linux GOARCH=amd64 go build -o build/${PROJECT_NAME}-linux-x64
echo "finished building linux version"

echo "building windows-x64 version"
env GOOS=windows GOARCH=amd64 go build -o build/${PROJECT_NAME}-windows-x64
echo "finished building windows version"

echo "building mac-x64 version"
env GOOS=darwin GOARCH=amd64 go build -o build/${PROJECT_NAME}-mac-x64
echo "finished building mac version"

echo "building linux-x86 version"
env GOOS=linux GOARCH=386 go build -o build/${PROJECT_NAME}-linux-x86
echo "finished building linux version"

echo "building windows-x86 version"
env GOOS=windows GOARCH=386 go build -o build/${PROJECT_NAME}-windows-x86
echo "finished building windows version"

echo "building mac-x86 version"
env GOOS=darwin GOARCH=386 go build -o build/${PROJECT_NAME}-mac-x86
echo "finished building mac version"

echo "copying config.toml and input.txt"
cp config.toml build/config.toml
cp input.txt build/input.txt

echo "build finished"