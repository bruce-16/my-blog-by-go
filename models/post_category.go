package models

import (
	"database/sql"
	"log"

	db "github.com/zachrey/blog/database"
)

type MPostAndCategory struct {
	Id         int64 `xorm:"pk autoincr"`
	PostId     int64 `xorm:"'post_id'"`
	CategoryId int64 `xorm:"'category_id'"`
	CreateTime int64 `xorm:"created 'create_time'"`
}

type MCategoryAndPost struct {
	CategoryId int64 `xorm:"'category_id'"`
	PostId     int64 `xorm:"'post_id'"`
	MPost      `xorm:"extends"`
}

// InsertPostAndCategory 将标题插入到post_category表
func InsertPostAndCategory(PostId, CategoryId int64) (int64, error) {
	mPostAndCategory := &MPostAndCategory{
		PostId:     PostId,
		CategoryId: CategoryId,
	}
	_, err := db.ORM.Table("post_category").Insert(mPostAndCategory)
	return mPostAndCategory.Id, err
}

// GetPostsByPLId 根据该表里面的id
func GetPostsByPCId(categoryId int64) *[]MCategoryAndPost {
	posts := make([]MCategoryAndPost, 0)
	err := db.ORM.
		Table("post_category").
		Join("INNER", "posts", "post_category.post_id=posts.id").
		Where("post_category.category_id=?", categoryId).
		Find(&posts)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	return &posts
}

// RemovePCByPostID 根据postid删除对于的记录
func RemovePCByPostID(postID int64) (sql.Result, error) {
	sql := "DELETE FROM post_category WHERE post_id=?"
	affacted, err := db.ORM.Sql(sql, postID).Execute()
	return affacted, err
}
