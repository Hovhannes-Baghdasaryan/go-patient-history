package config

import (
	_ "github.com/flashlabs/rootpath"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"

	"os"
)

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

type MainConfig struct {
	Env string `yaml:"env" env-default:"local"`
	HTTPServer
}

func MainConfigLoad() *MainConfig {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return nil
	}

	// check if file exist
	if _, err := os.Stat("config/main/config.yml"); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist")
		return nil
	}

	var cfg MainConfig
	if err := cleanenv.ReadConfig("config/main/config.yml", &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
		return nil
	}

	return &cfg
}
