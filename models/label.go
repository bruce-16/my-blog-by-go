package models

import (
	"log"
	"regexp"

	db "github.com/zachrey/blog/database"
)

type MLabel struct {
	Id         int64  `xorm:"pk autoincr"`
	Label      string `xorm:"unique 'label'"`
	CreateTime int64  `xorm:"created 'create_time'"`
}

// GetLabels 获取所有的标签
func GetLabels() *[]MLabel {
	var labels []MLabel
	err := db.ORM.Table("labels").Find(&labels)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	return &labels
}

func InsertLabel(labels string, ch chan []int64) {
	labelsArr := regexp.MustCompile(`\s*,\s*`).Split(labels, -1)
	ids := make([]int64, len(labelsArr))
	for i, v := range labelsArr {
		newLabel := &MLabel{Label: v}
		db.ORM.Table("labels").Insert(newLabel)
		if newLabel.Id == 0 {
			var newLabel2 MLabel
			db.ORM.Table("labels").Where("labels.label=?", v).Get(&newLabel2)
			ids[i] = newLabel2.Id
		} else {
			ids[i] = newLabel.Id
		}
	}
	ch <- ids
}
