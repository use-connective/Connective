package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var Config Configuration

type Configuration struct {
	Database           string   `json:"database"`
	DatabaseName       string   `json:"databaseName"`
	DatabaseUser       string   `json:"databaseUser"`
	DatabaseHost       string   `json:"databaseHost"`
	DatabasePort       string   `json:"databasePort"`
	DatabasePassword   string   `json:"databasePassword"`
	DatabaseSSLMode    string   `json:"databaseSSLMode"`
	JWTSecret          string   `json:"jwtSecret"`
	BackendPort        string   `json:"backendPort"`
	Environment        string   `json:"environment"`
	CORSAllowedOrigins []string `json:"corsAllowedOrigins"`
	RedisAddr          string   `json:"redisAddr"`
	RedisPassword      string   `json:"redisPassword"`
	RedisDB            int      `json:"redisDB"`
	RedisPoolSize      int      `json:"redisPoolSize"`
	RedisMinIdleConns  int      `json:"redisMinIdleConns"`
}

func LoadConfiguration() {
	configName := "config.local.json"
	env := os.Getenv("ENV")

	if env == "prod" {
		configName = "config.prod.json"
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("unable to read config file.")
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to unmarshal config file.")
	}
}
