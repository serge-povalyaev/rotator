package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	GRPC   GRPC
	DB     DBConfig
	Logger LoggerConfig
	Rabbit RabbitConfig
}

type GRPC struct {
	Host string
	Port string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func (c DBConfig) CreateDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Pass, c.Host, c.Port, c.Name)
}

type LoggerConfig struct {
	Level    string
	FilePath string
}

type RabbitConfig struct {
	Host  string
	Port  string
	User  string
	Pass  string
	Queue string
}

func (c RabbitConfig) CreateDSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s", c.User, c.Pass, c.Host, c.Port)
}

func ReadConfig(configPath string) Config {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Не удалось прочитать файл конфигурации: %s", err))
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Не удалось получить конфигурацию: %s", err))
	}

	return config
}
