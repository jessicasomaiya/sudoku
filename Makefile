MAIN_DIR = cmd/makeboard

.PHONY: local mpi

local: 
	go run $(MAIN_DIR)/main.go 

api: 
	go run $(MAIN_DIR)/main.go -runtime=api