package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var MainConfig *Config

type Config struct {
	DB struct {
		Host    string `json:"host"`
		Port    string `json:"port"`
		User    string `json:"user"`
		Pass    string `json:"pass"`
		Name    string `json:"name"`
		SSLMode string `json:"sslmode"`
	} `json:"db"`
	Location struct {
		Jakarta *time.Location
	}
}

func GetConfig() *Config {
	return MainConfig
}

func init() {
	var byteValue []byte
	jsonFile, err := os.Open("./config/local.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &MainConfig)
	MainConfig.Location.Jakarta, _ = time.LoadLocation("Asia/Jakarta")
}
