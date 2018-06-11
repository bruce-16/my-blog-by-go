package models

import (
	"database/sql"
	"log"

	db "github.com/zachrey/blog/database"
)

type MPost struct {
	Id         int64  `xorm:"pk autoincr"`
	Title      string `xorm:"'title'"`
	FileName   string `xorm: "file_name"`
	TextAmount int64  `xorm: "text_amount"`
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

// GetPostByTitle 根据标题获取post
func GetPostByTitle(title string) *MPost {
	var post MPost
	has, err := db.ORM.Table("posts").Where("title=?", title).Get(&post)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	if has == false {
		return nil
	}
	return &post
}

// GetPosts 获取所有的文章
func GetPosts() *[]MPost {
	var post []MPost
	err := db.ORM.Table("posts").Find(&post)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	return &post
}

// InsertPost 将标题插入到posts表
func InsertPost(title, fileName string, textAmount int64, ch chan int64) {
	newPost := new(MPost)
	newPost.Title = title
	newPost.FileName = fileName
	newPost.TextAmount = textAmount
	db.ORM.Table("posts").Insert(newPost)
	ch <- newPost.Id
}

// RemovePostByID 根据ID删除post
func RemovePostByID(ID int64) (sql.Result, error) {
	sql := "DELETE FROM posts WHERE id=?"
	affacted, err := db.ORM.Sql(sql, ID).Execute()
	return affacted, err
}
