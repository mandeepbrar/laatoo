#!/bin/sh
go build server/laatoo.go
rm -fr dist/bin/
mkdir -p dist/bin
cp --parents laatoo dist/bin/
