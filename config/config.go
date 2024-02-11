package config

import (
	"flag"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Sever    Server   `mapstructure:"server"`
	Logger   Logger   `mapstructure:"logger"`
	Database Database `mapstructure:"database"`
}

type Database struct {
	Server    string `mapstructure:"server"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	SSLMode   string `mapstructure:"sslmode"`
	Migration bool   `mapstructure:"migration"`
}
type Server struct {
	Name         string `mapstructure:"name"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	WriteTimeout int64  `mapstructure:"writetimeout"`
	ReadTimeout  int64  `mapstructure:"readtimeout"`
}
type Logger struct {
	WriteToFile bool   `mapstructure:"writeToFile"`
	Format      string `mapstructure:"format"`
}

func InitConfigs() *Config {
	path := fetchConfigPath()

	if path == "" {
		log.Fatalf("config path is empty %s", path)
	}

	var config Config

	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read configs %s", err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("failed to unmarshal configs %s", err.Error())
	}

	return &config
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
