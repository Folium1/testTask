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
	// Start the server
	delivery.StartServer()
}
