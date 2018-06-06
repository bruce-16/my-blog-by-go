package models

import (
	db "github.com/zachrey/blog/database"
)

type MPostAndCategory struct {
	Id         int64 `xorm:"pk autoincr"`
	PostId     int64 `xorm:"'post_id'"`
	CategoryId int64 `xorm:"'category_id'"`
	CreateTime int64 `xorm:"created 'create_time'"`
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
