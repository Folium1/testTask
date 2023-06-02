package main

import (
	"test/config"
	"test/delivery"
)

func main() {
	// Initialize the tables
	err := config.TablesInit()
	if err != nil {
		panic(err)
	}
	// Create mock user
	config.CreateMockUser()
	// Create mock logs
	config.CreateMockLogs()
	// Start the server
	delivery.StartServer()
}
