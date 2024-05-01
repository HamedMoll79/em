package postgresql

import "fmt"

type Config struct {
	Username string `env:"POSTGRES_USERNAME,notEmpty"`
	Password string `env:"POSTGRES_PASSWORD,notEmpty"`
	Port     int    `env:"POSTGRES_PORT,notEmpty"`
	Host     string `env:"POSTGRES_HOST,notEmpty"`
	DBName   string `env:"POSTGRES_DBNAME,notEmpty"`
	Driver   string `env:"POSTGRES_DRIVER"`
	Schema   string `env:"POSTGRES_SCHEMA"`
}

const Public = "public"

func (config *Config) DSN() string {
	if config.Schema == "" {
		config.Schema = Public
	}

	return fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.DBName)
}