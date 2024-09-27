package configs

import (
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"os"
	"sync"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string

	AppPort        string
	SideServiceUrl string
}

var (
	config Config
	once   sync.Once
)

// create new config
func New(logger logger.Logger) *Config {
	once.Do(func() {
		config = Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			DBName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),

			AppPort:        os.Getenv("APP_PORT"),
			SideServiceUrl: os.Getenv("SIDE_SERVICE_URL"),
		}
	})
	logger.Info("Config", "Config init")
	return &config
}
