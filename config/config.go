// Package config loads and stores application configuration from an ini file.
package config

import (
	"database/sql"
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	SQLDriver string
	DBName    string
	LogFile   string
}

var DB *sql.DB
var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DBName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
