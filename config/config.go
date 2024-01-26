package config

import (
	_ "github.com/flashlabs/rootpath"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	HTTPServer
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

func ConfigLoad() *Config {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// check if file exist
	if _, err := os.Stat("config/config.yml"); os.IsNotExist(err) {
		log.Fatal("config file does not exist")
	}

	var cfg Config
	if err := cleanenv.ReadConfig("config/config.yml", &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
