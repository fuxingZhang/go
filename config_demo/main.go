package main

import (
	"config_demo/config"
	"fmt"
)

func main() {
	fmt.Println(config.RunMode)
	fmt.Println(config.Port)

	conf := config.LoadConf()
	fmt.Println(conf)
	fmt.Println(conf.RunMode)
	fmt.Println(conf.Port)
}
