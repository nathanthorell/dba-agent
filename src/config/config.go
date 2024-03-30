package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	DBServers []ServerConfig `toml:"db_servers"`
	DBQueries []QueryConfig  `toml:"db_queries"`
	Scheduler SchedulerConfig
}

type ServerConfig struct {
	Name                string `toml:"name"`
	ConnectionStringVar string `toml:"conn_string_variable"`
}

type QueryConfig struct {
	Name     string `toml:"name"`
	Query    string `toml:"query"`
	Schedule string `toml:"schedule"`
	Server   string `toml:"server"`
}

type SchedulerConfig struct {
	Interval string `toml:"interval"`
}

func LoadConfig(filePath string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
