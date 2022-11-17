package config

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Address     string
	Ports       string
	Port        string
	DB_Driver   string
	DB_Name     string
	DB_Address  string
	DB_Port     string
	DB_Username string
	DB_Password string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = InitConfig()
	}

	return appConfig
}

func InitConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Address = getEnv("ADDRESS", "http://localhost")
	defaultConfig.Port = getEnv("PORT", "3030")
	defaultConfig.DB_Name = getEnv("MYSQL_DBNAME", "todolist-simple")
	defaultConfig.DB_Address = getEnv("MYSQL_HOST", "172.17.0.1")
	defaultConfig.DB_Port = getEnv("MYSQL_PORT", "3306")
	defaultConfig.DB_Username = getEnv("MYSQL_USER", "root")
	defaultConfig.DB_Password = getEnv("MYSQL_PASSWORD", "secret")

	fmt.Println(defaultConfig)
	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}

	return fallback
}
