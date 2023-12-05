package main

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Text string `json:"text"`
}
