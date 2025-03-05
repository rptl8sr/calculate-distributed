package config

import (
	"fmt"
	"log/slog"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// App
	AppLogLevel   slog.Level `envconfig:"APP_LOG_LEVEL" default:"info"`
	AppPort       int        `envconfig:"APP_PORT" default:"8080"`
	MaxGoroutines int        `envconfig:"MAX_GOROUTINES" default:"10"`
}

func Must() *Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		panic(fmt.Sprintf("Error processing environment variables: %v", err))
	}

	return &config
}
