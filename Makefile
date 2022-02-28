MAIN_DIR = packages/api/cmd

all: api

api:
	@ go run $(MAIN_DIR)/main.go 

local: test
	@ go run $(MAIN_DIR)/main.go  -local

test: 
	@go test -v ./...
