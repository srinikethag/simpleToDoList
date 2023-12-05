package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func initDB() {
	db, err = gorm.Open("postgres", "host=localhost user=youruser dbname=yourdb sslmode=disable password=yourpassword")
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Todo{})
}
