package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	App struct {
		Port string `yaml:"port" env-default:"9000"`
		Host string `yaml:"host" env-default:"localhost"`
	} `yaml:"app"`

	Database struct {
		Port     string `yaml:"port" env-default:"5432"`
		Host     string `yaml:"host" env-default:"localhost"`
		Name     string `yaml:"name" env-default:"postgres"`
		User     string `yaml:"user" env-default:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

func LoadEnv() *Config {
	var cfg Config

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
