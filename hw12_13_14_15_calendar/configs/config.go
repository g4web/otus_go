package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	LogFile     string `mapstructure:"LOG_FILE"`
	StorageType string `mapstructure:"STORAGE_TYPE"`
	DbHost      string `mapstructure:"DB_HOST"`
	DbPort      string `mapstructure:"DB_PORT"`
	DbName      string `mapstructure:"DB_NAME"`
	DbUser      string `mapstructure:"DB_USER"`
	DbPassword  string `mapstructure:"DB_PASSWORD"`
	HttpHost    string `mapstructure:"HTTP_HOST"`
	HttpPort    string `mapstructure:"HTTP_PORT"`
	GrpcHost    string `mapstructure:"GRPC_HOST"`
	GrpcPort    string `mapstructure:"GRPC_PORT"`
}

func NewConfig(filePath string) (*Config, error) {
	var conf Config

	viper.SetConfigFile(filePath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
