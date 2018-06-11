package controllers

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/zachrey/blog/models"
)

var wg sync.WaitGroup

// RemoveFile 根据标题删除文章及相关项
func RemoveFile(c *gin.Context) {
	wg.Add(3)
	fileName := c.Query("name")
	if fileName == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "Filed name is empty.",
		})
		return
	}

	post := models.GetPostByTitle(fileName)

	// 删除分类关联表
	go func() {
		models.RemovePCByPostID(post.Id)
		wg.Done()
	}()
	// 删除标签关联表
	go func() {
		models.RemovePLByPostID(post.Id)
		wg.Done()
	}()
	// 删除文件
	go func() {
		err := os.Remove("./posts/" + post.FileName)
		if err != nil {
			log.Println("[ERROR] ", err)
		}
		wg.Done()
	}()
	wg.Wait()
	models.RemovePostByID(post.Id)

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "删除文章成功!",
	})
}
