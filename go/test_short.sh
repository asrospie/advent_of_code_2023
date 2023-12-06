#!/bin/bash

go clean -testcache
go test ./... -short
