package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration config struct
type Configuration struct {
	Port    string `json:"port"`
	RunMode string `json:"run_mode"`
}

// LoadConf load json config file
func LoadConf() (conf Configuration) {
	file, err := os.Open("./config/conf.json")
	defer file.Close()
	if err != nil {
		fmt.Println("read json file error:", err)
	}
	decoder := json.NewDecoder(file)
	e := decoder.Decode(&conf)
	if e != nil {
		fmt.Println("parse json error:", e)
	}
	return conf
}
