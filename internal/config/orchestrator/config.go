package config

import (
	"fmt"
	"log/slog"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// App
	AppLogLevel slog.Level `envconfig:"APP_LOG_LEVEL" default:"info"`
	AppPort     int        `envconfig:"APP_PORT" default:"8080"`
	Timeouts    Timeouts
}

type Timeouts struct {
	Addition       int `envconfig:"TIME_ADDITION_MS" default:"1000"`
	Subtraction    int `envconfig:"TIME_SUBTRACTION_MS" default:"1000"`
	Multiplication int `envconfig:"TIME_MULTIPLICATION_MS" default:"1000"`
	Division       int `envconfig:"TIME_DIVISION_MS" default:"1000"`
}

func Must() *Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		panic(fmt.Sprintf("Error processing environment variables: %v", err))
	}

	return &config
}
