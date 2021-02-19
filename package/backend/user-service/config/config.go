package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Port, DBHost, DBDebug, DBDriver, DBUser, DBPassword, DBName, DBPort string
}

var config Config
var env = "dev"
var environments = map[string]string{
	"dev":        "./dev.yaml",
	"production": "./prod.yaml",
}

func Init() {
	LoadConfigByEnv("dev")
}

func LoadConfigByEnv(env string) {
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
