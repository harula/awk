package main

import (
	"sync"
)

type Config struct {
	name string
	age  int
}

var once sync.Once
var config *Config

func loadConfig() *Config {
	var cfg Config
	cfg.name = "test"
	cfg.age = 10

	return &cfg
}

func initialize() {
	config = loadConfig()
}

func getConfig() *Config {
	once.Do(initialize) //once.Do适合再读配置，建立数据库连接等场景
	return config
}
