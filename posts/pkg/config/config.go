package config

import (
	"github.com/spf13/viper"
	"time"
)

// Config stores all configurations of the applications
// The values are read by viper from a config file or environment variables
type Config struct {
	MONGOURI             string        `mapstructure:"MONGO_URI"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads the data fromm the config file
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)

	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
