package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Port, Secret, RedisPort, RedisHost, UserAddr string
}

var config Config
var env = "dev"
var environments = map[string]string{
	"dev":        "./dev.yaml",
	"production": "./prod.yaml",
}

func Init() {
	LoadSettingsByEnv("dev")
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	config = Config{}
	jsonErr := json.Unmarshal(content, &config)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Config {
	if &config == nil {
		Init()
	}
	return config
}
