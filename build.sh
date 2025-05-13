#!/bin/zsh
set -e

mkdir -p build
go build -o build/go-webscraper
echo "✅ Built at build/go-webscraper"