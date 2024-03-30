package scheduler

import (
	"fmt"
	"log"
	"os"

	"github.com/nathanthorell/dba-agent/src/config"
	"github.com/nathanthorell/dba-agent/src/db"
	"github.com/robfig/cron/v3"
)

// Scheduler struct represents the scheduler
type Scheduler struct {
	config    config.Config
	scheduler *cron.Cron
}

// NewScheduler creates a new scheduler
func NewScheduler(config config.Config) *Scheduler {
	return &Scheduler{
		config:    config,
		scheduler: cron.New(),
	}
}

func (s *Scheduler) getServerByName(serverName string) (config.ServerConfig, error) {
	for _, server := range s.config.DBServers {
		if server.Name == serverName {
			return server, nil
		}
	}
	return config.ServerConfig{}, fmt.Errorf("server not found: %s", serverName)
}

func (s *Scheduler) Start() {
	log.Printf("Scheduler started")

	for _, query := range s.config.DBQueries {
		// Get server configuration
		server, err := s.getServerByName(query.Server)
		if err != nil {
			log.Printf("Error getting server config for query %s: %v", query.Name, err)
			continue
		}

		// Schedule task to execute query
		_, err = s.scheduler.AddFunc(query.Schedule, func() {
			connectionString := os.Getenv(server.ConnectionStringVar)
			if connectionString == "" {
				log.Printf("Empty connection string for server %s", server.Name)
				return
			}

			db, err := db.NewDBConnection(connectionString)
			if err != nil {
				log.Printf("Error connecting to database for query %s: %v", query.Name, err)
				return
			}
			defer db.Close()

			// Execute query
			_, err = db.Exec(query.Query)
			if err != nil {
				log.Printf("Error executing query %s: %v", query.Name, err)
			}
		})
		if err != nil {
			log.Printf("Error scheduling task for query %s: %v", query.Name, err)
		}
	}

	s.scheduler.Start()
}
