MAIN_DIR = cmd

local: 
	go run $(MAIN_DIR)/local/main.go 

api: 
	go run $(MAIN_DIR)/api/main.go