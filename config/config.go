package config

import "fmt"

const DefaultConfigPath = "/Users/justintieri/.habbgo/config/config.yaml"

// Config represents the entire set of configuration options for the application.
type Config struct {
	Global *GlobalConfig
	Server *ServerConfig
	DB     *DatabaseConfig
}

// GlobalConfig represents the configuration options that are applied globally across the application.
type GlobalConfig struct {
	Debug bool `yaml:"debug" json:"debug"`
}

// ServerConfig represents the configuration options for the underlying tcp game server connection.
type ServerConfig struct {
	Host              string `yaml:"host" json:"host"`
	Port              int    `yaml:"port" json:"port"`
	MaxConnsPerPlayer int    `yaml:"max-connections-per-player" json:"max-connections-per-player"`
}

// DatabaseConfig represents the configuration options for the underlying database connection.
type DatabaseConfig struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Name     string `yaml:"db-name" json:"db-name"`
	Driver   string `yaml:"driver" json:"driver"`
	SSLMode  string `yaml:"ssl-mode" json:"ssl-mode"`
}

// ConnectionString returns a database connection string for Postgres.
func (d *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		d.Username, d.Password, d.Host, d.Port, d.Name, d.SSLMode)
}

// defaultConfig returns a default configuration object to be marshaled to disk.
func defaultConfig() *Config {
	return &Config{
		Global: &GlobalConfig{
			Debug: false,
		},
		Server: &ServerConfig{
			Host:              "127.0.0.1",
			Port:              11235,
			MaxConnsPerPlayer: 2,
		},
		DB: &DatabaseConfig{
			Username: "anon",
			Password: "password",
			Host:     "127.0.0.1",
			Port:     5432,
			Name:     "habbgo",
			Driver:   "postgres",
			SSLMode:  "disable",
		},
	}
}
