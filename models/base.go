// Package models provides database models and initialization logic.
package models

import (
	"go-echo-simple-example/config"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const (
	tableNameTodo = "todos"
)

func init() {
	if config.Config.DbName == "" {
		log.Fatal("database name is not set in config")
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(config.Config.DbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := DB.AutoMigrate(&Todo{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
