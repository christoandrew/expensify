package main

import (
	"expensify/db"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func getDB() (*gorm.DB, error) {
	b, err := gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	return b, err
}

func GetDB() *gorm.DB {
	return DB
}

func InitDB() {
	d, err := getDB()
	if err != nil {
		log.Fatal(err)
	} else {
		DB = d
	}
}

func Migrate(registry db.ModelRegistry) {
	for _, model := range registry.GetModels() {
		DB.AutoMigrate(model)
	}
}

func CloseDB() {
	err := DB.Close()
	CheckError(err)
}
