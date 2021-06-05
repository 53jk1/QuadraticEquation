#!make
SHELL := /bin/bash

.PHONY: run
run: ##@run Runs QuadraticEquation
	go run main.go

.PHONY: test
test: ##@test Runs tests all services
	go test -cover -v