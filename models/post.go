package models

import (
	"log"

	db "github.com/zachrey/blog/database"
)

type MPost struct {
	Id         int64  `xorm:"pk autoincr"`
	Title      string `xorm:"'title'"`
	FileName   string `xorm: "file_name"`
	CreateTime int64  `xorm:"created 'create_time'"`
}

// GetPostByID 根据ID获取文章
func GetPostByID(Id int64) *MPost {
	var post MPost
	has, err := db.ORM.Table("posts").Id(Id).Get(&post)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	if has == false {
		return nil
	}
	return &post
}

// InsertPost 将标题插入到posts表
func InsertPost(title, fileName string, ch chan int64) {
	newPost := new(MPost)
	newPost.Title = title
	newPost.FileName = fileName
	db.ORM.Table("posts").Insert(newPost)
	ch <- newPost.Id
}
