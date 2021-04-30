#!/bin/sh -eux
go test main_test.go main.go -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
open cover.html
