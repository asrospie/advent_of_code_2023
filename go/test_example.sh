#!/bin/bash

go clean -testcache
go test ./... -run TestDay5Part2
