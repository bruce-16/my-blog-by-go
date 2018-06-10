package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zachrey/blog/models"
)

//GetLabels 获取所有的标签
func GetLabels(c *gin.Context) {
	labels := models.GetLabels()
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   labels,
	})
}
