package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// struct tags bolte hai na ki anotations := `ymal:"env" env:"ENV" env-required:"true" env-default:"production"`
type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `ymal:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `ymal:"storage_path" env-required:"true"`
	HTTPServer  `ymal:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("no config path provided")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("failed to read config: %s", err.Error())
	}

	return &cfg

}
