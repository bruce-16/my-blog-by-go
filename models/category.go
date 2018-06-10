package models

import (
	"log"
	"regexp"

	db "github.com/zachrey/blog/database"
)

type MCategory struct {
	Id         int64  `xorm:"pk autoincr"`
	Category   string `xorm:"unique 'category'"`
	CreateTime int64  `xorm:"created 'create_time'"`
}

// GetLabels 获取所有的标签
func GetCategories() *[]MCategory {
	var categories []MCategory
	err := db.ORM.Table("categories").Find(&categories)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	return &categories
}

func InsertCategory(categories string, ch chan []int64) {
	categoriesArr := regexp.MustCompile(`\s*,\s*`).Split(categories, -1)
	ids := make([]int64, len(categoriesArr))
	for i, v := range categoriesArr {
		newCategories := &MCategory{Category: v}
		db.ORM.Table("categories").Insert(newCategories)
		if newCategories.Id == 0 {
			var newCategories2 MCategory
			db.ORM.Table("categories").Where("categories.category=?", v).Get(&newCategories2)
			ids[i] = newCategories2.Id
		} else {
			ids[i] = newCategories.Id
		}
	}
	ch <- ids
}
