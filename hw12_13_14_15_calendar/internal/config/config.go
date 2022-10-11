package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel        string `mapstructure:"LOG_LEVEL"`
	LogFile         string `mapstructure:"LOG_FILE"`
	StorageType     string `mapstructure:"STORAGE_TYPE"`
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          string `mapstructure:"DB_PORT"`
	DBName          string `mapstructure:"DB_NAME"`
	DBUser          string `mapstructure:"DB_USER"`
	DBPassword      string `mapstructure:"DB_PASSWORD"`
	HTTPHost        string `mapstructure:"HTTP_HOST"`
	HTTPPort        string `mapstructure:"HTTP_PORT"`
	GRPCHost        string `mapstructure:"GRPC_HOST"`
	GRPCPort        string `mapstructure:"GRPC_PORT"`
	SchedulerPeriod string `mapstructure:"SCHEDULER_PERIOD"`
	MQAddr          string `mapstructure:"MQ_ADDR"`
	MQQueue         string `mapstructure:"MQ_QUEUE"`
	MQHandlersCount int    `mapstructure:"MQ_HANDLERS_COUNT"`
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
