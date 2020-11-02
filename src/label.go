package main

import (
	"gorm.io/gorm"
	"log"
)

type Label struct {
	gorm.Model
	name string `gorm:"unique"`
}

func GetAllLabels() *[]Label {
	conn := GetConnection()
	var labels []Label
	if result := conn.db.Table("labels").Find(&labels); result.Error != nil {
		log.Println("getAllLabels Error:", result.Error)
		return nil
	}
	return &labels
}

func InsertLabel(label string) uint {
	newLabel := &Label{name: label}
	conn := GetConnection()
	conn.db.Table("categories").Create(newLabel)
	return newLabel.ID
}
