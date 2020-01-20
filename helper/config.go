package helper

import (
	"log"

	"github.com/BurntSushi/toml"
)

type config struct {
	username,
	password,
	session string
}

func getConfig() *config {
	c := new(config)
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}
