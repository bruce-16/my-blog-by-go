package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zachrey/blog/models"
)

//GetPosts 获取所有的文章
func GetPosts(c *gin.Context) {
	labels := models.GetPosts()
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   labels,
	})
}

//GetPostByLabelId 根据label id获取post
func GetPostByLabelId(c *gin.Context) {
	labelid := c.Param("labelid")
	if labelid == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "labelid 不能为空",
		})
	}
	labelId, err := strconv.ParseInt(labelid, 10, 64)
	posts := models.GetPostsByPLId(labelId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "labelid can't convert to int64, the error information: " + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   posts,
	})
}

//GetPostByCategoryId category id获取post
func GetPostByCategoryId(c *gin.Context) {
	categoryid := c.Param("categoryid")
	if categoryid == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "categoryid 不能为空",
		})
	}
	categoryId, err := strconv.ParseInt(categoryid, 10, 64)
	posts := models.GetPostsByPCId(categoryId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"status": 1,
			"msg":    "categoryid can't convert to int64, the error information: " + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   posts,
	})
}
