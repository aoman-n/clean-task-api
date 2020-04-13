package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server ServerConfig
	GoEnv  string
}

type ServerConfig struct {
	Port string
}

var conf Config

func Get() Config {
	return conf
}

func init() {
	env := os.Getenv("GO_ENV")
	filepath := fmt.Sprintf("./src/config/%v.toml", env)
	conf = Config{GoEnv: env}

	if _, err := toml.DecodeFile(filepath, &conf); err != nil {
		log.Fatal("failed to load config toml ", err)
	}
}
