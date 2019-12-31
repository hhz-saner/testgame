package config

import (
	"flag"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Env         string
	Http        Http
	Database    Database
	Cache       Cache
}

type Http struct {
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

type Database struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Prefix   string
}

type Cache struct {
	Type     string
	Host     string
	Port     string
	Password string
	DB       int
}

var (
	Cfg Config
)

func init() {
	filePath := flag.String("config", "./config.toml", "file path")
	flag.Parse()
	if _, err := toml.DecodeFile(*filePath, &Cfg); err != nil {
		panic(err.Error())
	}
}