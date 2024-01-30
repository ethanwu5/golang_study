package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}

type Mysql struct {
	Url  string
	Port int
}

type Redis struct {
	Host string
	Port int
}

func main() {
	dataBytes, err := os.ReadFile("./conf.yaml")
	if err != nil {
		panic(err)
	}
	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		panic(err)
	}

	fmt.Println(config.Mysql.Url)
	fmt.Println(config.Redis.Host)
}
