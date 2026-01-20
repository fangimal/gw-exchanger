package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPC struct {
		Port string `yaml:"port" env-default:"50051"`
	} `yaml:"grpc"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	DBHost     string `yaml:"host" env-default:"localhost"`
	DBPort     string `yaml:"port" env-default:"5432"`
	DBUser     string `yaml:"username" env-default:"wallet_user"`
	DBPassword string `yaml:"password" env-default:"123"`
	DBName     string `yaml:"database" env-default:"wallet_db"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			fmt.Printf("Error reading config: %v", help)
		}
	})
	return instance
}
