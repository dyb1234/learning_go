package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	Posts     []Post
	PostCount uint
}
type Post struct {
	ID            uint `gorm:"primarykey"`
	Comments      []Comment
	CommentCount  uint
	CommentStatus string
	UserID        uint
}
type Comment struct {
	ID     uint `gorm:"primarykey"`
	PostID uint
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(&User{}).Where("ID = ?", post.UserID).Update("post_count", gorm.Expr("post_count + ?", 1)).Error
	return err
}

func (co *Comment) AfterDelete(tx *gorm.DB) (err error) {
	postID := co.PostID
	tx.Model(&Post{}).Where("ID = ?", postID).Update("comment_count", gorm.Expr("comment_count - ?", 1))
	var count uint
	err = tx.Model(&Post{}).Select("comment_count").Where("ID = ?", postID).First(&count).Error
	if count == 0 {
		tx.Model(&Post{}).Where("ID = ?", postID).Update("comment_status", "无评论")
	}
	return err
}
func main() {
	dsn := "root:root@tcp(127.0.0.1:13306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	// create tables
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		return
	}
	db = db.Debug()
	// question 3.2
	var user User
	db.Preload("Posts.Comments").Where("users.id = ?", 1).Find(&user)
	jsonData, _ := json.MarshalIndent(user, "", " ")
	fmt.Println(string(jsonData))
	var post Post
	db.Preload("Comments").Order("comment_count DESC").First(&post)
	jsonData, _ = json.MarshalIndent(post, "", " ")
	fmt.Println(string(jsonData))
	// question 3.3
	//postNew := Post{CommentStatus: "无评论", CommentCount: 0, UserID: 1}
	//db.Create(&postNew)
	comment := Comment{1, 1}
	db.Where("ID = ?", 1).Delete(&comment)
}
