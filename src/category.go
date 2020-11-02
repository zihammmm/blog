package main

import (
	"gorm.io/gorm"
	"log"
)

type Category struct {
	gorm.Model
	name string
}

func GetAllCategories() *[]Category {
	conn := GetConnection()
	var categories []Category
	if result := conn.db.Table("categories").Find(&categories); result.Error != nil {
		log.Println("GetAllCategories Error:", result.Error)
		return nil
	}
	return &categories
}

func InsertCategory(category string) uint {
	newCategory := &Category{name: category}
	conn := GetConnection()
	conn.db.Table("categories").Create(newCategory)
	return newCategory.ID
}
