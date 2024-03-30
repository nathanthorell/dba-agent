package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nathanthorell/dba-agent/src/config"
	"github.com/nathanthorell/dba-agent/src/scheduler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Initialize configuration
	cfg, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Start scheduler
	sched := scheduler.NewScheduler(*cfg)
	sched.Start()

	// Keep the main thread running
	select {}
}
