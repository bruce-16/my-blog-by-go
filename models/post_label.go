package models

import (
	db "github.com/zachrey/blog/database"
)

type MPostAndLabel struct {
	Id         int64 `xorm:"pk autoincr"`
	PostId     int64 `xorm:"'post_id'"`
	LabelId    int64 `xorm:"'Label_id'"`
	CreateTime int64 `xorm:"created 'create_time'"`
}

// InsertPostAndLabel 将标题插入到post_label表
func InsertPostAndLabel(PostId, LabelId int64) (int64, error) {
	mPostAndLabel := &MPostAndLabel{
		PostId:  PostId,
		LabelId: LabelId,
	}
	_, err := db.ORM.Table("post_label").Insert(mPostAndLabel)
	return mPostAndLabel.Id, err
}
