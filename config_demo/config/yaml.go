package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// YamlConfig yaml config
type YamlConfig struct {
	Redis struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
	}
	Mysql struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

/** or */
// type Config struct{
// 	Redis Redis `yaml:"redis"` //或者：`yaml:"redis, inline"`
// 	Mysql Mysql `yaml:"mysql"` //或者：`yaml:"mysql, inline"`
// }
// type Redis struct{
// 	Address string `yaml:"address"`
// 	Password string `yaml:"password"`
// }
// type Mysql struct{
// 	Username string `yaml:"username"`
// 	Password string `yaml:"password"`
// }

// DatabaseConfig DatabaseConfig
var DatabaseConfig *YamlConfig

// LoadYaml LoadYaml
func LoadYaml(path string) *YamlConfig {
	conf := new(YamlConfig)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("yamlFile.Get err:", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatal("Unmarshal:", err)
	}

	return conf
}
