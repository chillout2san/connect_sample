#!/bin/bash

# パッケージを抽出する
packages=$(grep -oE "\"(.*)\"" ./tool/tool.go | tr -d '"' | tr '\n' ' ')

# それぞれを go install する
for package in $packages
do
    go install $package
done

echo "go install done"

go mod tidy

echo "go mod tidy done"