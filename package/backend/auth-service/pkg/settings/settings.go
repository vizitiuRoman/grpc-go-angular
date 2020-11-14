package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Settings struct {
	Port, Secret, RedisPort, RedisHost string
}

var settings Settings
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
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}
