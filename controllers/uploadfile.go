package controllers

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/zachrey/blog/models"
)

const postDir = "./posts/"

var gr sync.WaitGroup
var isShouldRemove = false

// UpLoadFile 上传文件的控制器
func UpLoadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	filename := header.Filename
	md5FileName := fmt.Sprintf("%x", md5.Sum([]byte(filename)))
	fileExt := filepath.Ext(postDir + filename)
	filePath := postDir + md5FileName + fileExt
	log.Println("[INFO] upload file: ", header.Filename)

	has := hasSameNameFile(md5FileName+fileExt, postDir)
	if has {
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"msg":    "服务器已有相同文件名称",
		})
		return
	}

	// 根据文件名的md5值，创建服务器上的文件
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 处理完整个上传过程后，是否需要删除创建的文件，在存在错误的情况下
	defer func() {
		if isShouldRemove {
			err = os.Remove(filePath)
			if err != nil {
				log.Println("[ERROR] ", err)
			}
		}
	}()
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	err = readMdFileInfo(filePath)
	if err != nil {
		isShouldRemove = true
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
			go models.InsertPost(mdInfo[TITLE], filepath.Base(filePath), postCh)
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
