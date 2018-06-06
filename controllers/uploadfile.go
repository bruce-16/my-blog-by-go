package controllers

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/zachrey/blog/models"
)

const postDir = "./posts/"

var gr sync.WaitGroup

// UpLoadFile 上传文件的控制器
func UpLoadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	filename := header.Filename
	log.Println("[INFO] upload file: ", header.Filename)
	has := hasSameNameFile(filename, postDir)
	if has {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "服务器已有相同文件名称",
		})
		return
	}
	out, err := os.Create(postDir + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	err = readMdFileInfo(postDir + filename)
	if err != nil {
		// out.Close()
		// err = os.RemoveAll(postDir + filename)
		// if err != nil {
		// 	log.Println("[ERROR] ", err)
		// }
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "上传成功",
	})
}

func readMdFileInfo(filePath string) error {
	fileread, _ := ioutil.ReadFile(filePath)
	lines := strings.Split(string(fileread), "\n")
	log.Println(lines)
	const (
		TITLE      = "title: "
		CATEGORIES = "categories: "
		LABEL      = "label: "
	)
	var (
		postId int64
	)
	mdInfo := make(map[string]string)
	var (
		postCh     chan int64
		categoryCh chan []int64
		labelCh    chan []int64
	)
	for i, lens := 0, len(lines); i < lens && i < 5; i++ { // 只查找前五行
		switch {
		case strings.HasPrefix(lines[i], TITLE):
			mdInfo[TITLE] = strings.TrimLeft(lines[i], TITLE)
			postCh = make(chan int64)
			go models.InsertPost(mdInfo[TITLE], postCh)
		case strings.HasPrefix(lines[i], CATEGORIES):
			mdInfo[CATEGORIES] = strings.TrimLeft(lines[i], CATEGORIES)
			categoryCh = make(chan []int64)
			go models.InsertCategory(mdInfo[CATEGORIES], categoryCh)
		case strings.HasPrefix(lines[i], LABEL):
			mdInfo[LABEL] = strings.TrimLeft(lines[i], LABEL)
			labelCh = make(chan []int64)
			go models.InsertLabel(mdInfo[LABEL], labelCh)
		}
	}
	postId = <-postCh
	if postId == 0 {
		return errors.New("服务器上已有相同文章标题")
	}
	log.Println("[INFO] postId: ", postId)
	if categoryCh != nil {
		go func() {
			categoryIds := <-categoryCh
			log.Println("[INFO] categoryIds: ", categoryIds)

			for _, v := range categoryIds {
				models.InsertPostAndCategory(v, postId)
			}
		}()
	}

	if labelCh != nil {
		go func() {
			labels := <-labelCh
			log.Println("[INFO] labels: ", labels)

			for _, v := range labels {
				models.InsertPostAndLabel(v, postId)
			}
		}()
	}
	return nil
	// 接下来就是将各种信息存入数据库
}

func hasSameNameFile(fileName, dir string) bool {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if fileName == f.Name() {
			return true
		}
	}
	return false
}
