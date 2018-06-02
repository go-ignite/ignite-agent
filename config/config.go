package config

import (
	"log"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	C Config
)

type Config struct {
	App struct {
		Address  string `mapstructure:"address"`
		LogLevel string `mapstructure:"log_level"`
	} `mapstructure:"app"`
	Host struct {
		Address string `mapstructure:"address"`
		From    int    `mapstructure:"from"`
		To      int    `mapstructure:"to"`
	} `mapstructure:"host"`
	Secret struct {
		User  string `mapstructure:"user"`
		Admin string `mapstructure:"admin"`
	} `mapstructure:"secret"`
}

func Init() {
	// app
	viper.SetDefault("app.log_level", "INFO")
	viper.SetDefault("app.address", ":4000")
	// host
	viper.SetDefault("host.address", "localhost")
	viper.SetDefault("host.from", "5001")
	viper.SetDefault("host.to", "6000")
	// secret
	viper.SetDefault("secret.user", "ignite-user")
	viper.SetDefault("secret.admin", "ignite-admin")

	// bind envs
	viper.SetEnvPrefix("agent")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&C); err != nil {
		log.Fatalf("viper.Unmarshal error: %v\n", err)
	}

	// log
	lv, err := logrus.ParseLevel(C.App.LogLevel)
	if err != nil {
		log.Fatalf("logrus.ParseLevel error: %v\n", err)
	}
	logrus.SetLevel(lv)
}
