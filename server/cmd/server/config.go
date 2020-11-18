package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/Alexander3006/architecture-lab-3/server/domain/interfaces"
)

const ConfigEnv = "LAB_CONFIG"
const DefaultConf = "config.json"

var once sync.Once
var instance *Config

type DBconfig struct {
	Name       string `json:"name"`
	User       string `json:"user"`
	Host       string `json:"host"`
	Password   string `json:"password"`
	DisableSSL bool   `json:"disableSSL"`
}

type ServerConfig struct {
	Port interfaces.HttpPortNumber `json:"port"`
}

type Config struct {
	Db     DBconfig     `json:"db"`
	Server ServerConfig `json:"server"`
}

func ReadConfig(path string) (*Config, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := Config{}
	if err := json.Unmarshal(dat, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func DefaultConfig() (*Config, error) {
	file, exists := os.LookupEnv(ConfigEnv)
	if !exists {
		file = DefaultConf
	}
	return ReadConfig(file)
}

func GetConfig() (*Config, error) {
	var err error = nil
	once.Do(func() {
		instance, err = DefaultConfig()
	})

	return instance, err
}

func GetPort(conf *Config) interfaces.HttpPortNumber {
	return conf.Server.Port
}
