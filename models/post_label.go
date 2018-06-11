package models

import (
	"database/sql"
	"log"

	db "github.com/zachrey/blog/database"
)

type MPostAndLabel struct {
	Id         int64 `xorm:"pk autoincr"`
	PostId     int64 `xorm:"'post_id'"`
	LabelId    int64 `xorm:"'Label_id'"`
	CreateTime int64 `xorm:"created 'create_time'"`
}

type LabelAndPost struct {
	LabelId int64 `xorm:"'label_id'"`
	PostId  int64 `xorm:"'post_id'"`
	MPost   `xorm:"extends"`
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

// GetPostsByPLId 根据该表里面的id
func GetPostsByPLId(labelId int64) *[]LabelAndPost {
	posts := make([]LabelAndPost, 0)
	err := db.ORM.
		Table("post_label").
		Join("INNER", "posts", "post_label.post_id=posts.id").
		Where("post_label.label_id=?", labelId).
		Find(&posts)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	return &posts
}

// RemoveByPostID 根据postid删除对于的记录
func RemovePLByPostID(postID int64) (sql.Result, error) {
	sql := "DELETE FROM post_label WHERE post_id=?"
	affacted, err := db.ORM.Sql(sql, postID).Execute()
	return affacted, err
}
