package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	LogLevel string `mapstructure:"log_level"`
	Address  string `mapstructure:"address"`
	Token    string `mapstructure:"token"`
}

var defaultConfig = Config{
	LogLevel: logrus.InfoLevel.String(),
	Address:  ":4000",
	Token:    "ignite-agent",
}

func Init() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, errors.Wrap(err, "config: load .env file error")
	}

	viper.SetEnvPrefix("agent")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	c := defaultConfig
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalf("viper.Unmarshal error: %v\n", err)
	}

	lv, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("config: log_level is invalid")
	}
	logrus.SetLevel(lv)

	return &c, nil
}
