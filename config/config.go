package config

import (
	"log"

	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env string
	Gin *GinConfig
	DB  *DBConfig
}

type GinConfig struct {
	Host            string
	Port            string
	Mode            string
	Domain          string
	Docs            string
	ContractVersion string
}

type DBConfig struct {
	Dialect  string
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Flag     string
}

func NewConfig() *Config {
	configPath := "./"
	runPath, _ := os.Getwd()
	matchPathStatus := false
	pathArr := strings.Split(runPath, "/")
	for i := len(pathArr) - 1; i > 0; i-- {
		configPath += "../"
		if pathArr[i] == "cmd" || pathArr[i] == "test" || pathArr[i] == "migration" {
			matchPathStatus = true
			break
		}
	}
	if !matchPathStatus {
		configPath = "./"
	}
	configPath += "config"

	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		Env: viper.GetString("env"),
		Gin: &GinConfig{
			Host:            viper.GetString("server.host"),
			Port:            viper.GetString("server.port"),
			Mode:            viper.GetString("server.mode"),
			Domain:          viper.GetString("server.domain"),
			Docs:            viper.GetString("server.docs"),
			ContractVersion: viper.GetString("server.contract_version"),
		},
		DB: &DBConfig{
			Dialect:  viper.GetString("db.dialect"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Database: viper.GetString("db.database"),
			Flag:     viper.GetString("db.flag"),
		},
	}
}
