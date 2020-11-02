package main

import (
	"gorm.io/gorm"
	"log"
)

type Post struct {
	gorm.Model
	title      string `gorm:"unique;not null"`
	filePath   string `gorm:"unique;not null"`
	labels     []uint
	Categories []uint
}

func GetPost() {

}

func GetPostByID(id uint) *Post {
	var post Post
	conn := GetConnection()
	if result := conn.db.Table("posts").Take(post); result.Error != nil {
		log.Println("GetPostByID Error:", result.Error)
		return nil
	}
	return &post
}

func GetPostByTitle(name string) *Post {
	var post Post
	conn := GetConnection()
	if result := conn.db.Table("posts").Where("name LIKE ?", "%"+name+"%").Find(&post); result.Error != nil {
		log.Println("GetPostByTitle Error:", result.Error)
		return nil
	}
	return &post
}

func GetAllPosts() *[]Post {
	var posts []Post
	conn := GetConnection()
	if result := conn.db.Table("posts").Find(&posts); result.Error != nil {
		log.Println("GetAllPost Error", result.Error)
		return nil
	}
	return &posts
}

func InsertPost(name string, filePath string) bool {
	return true
}

func RemovePostByID(ID uint) bool {
	return true
}
