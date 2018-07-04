package controllers

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"unicode"

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
	body := strings.Join(lines[5:], "")
	textAmount := GetStrLength(body)
	log.Println(lines)
	const (
		TITLE      = "title: "
		CATEGORIES = "categories: "
		LABEL      = "label: "
	)
	var (
		postId     int64
		postCh     chan int64
		categoryCh chan []int64
		labelCh    chan []int64
	)
	mdInfo := make(map[string]string)
	
	for i, lens := 0, len(lines); i < lens && i < 5; i++ { // 只查找前五行
		switch {
		case strings.HasPrefix(lines[i], TITLE):
			mdInfo[TITLE] = strings.TrimLeft(lines[i], TITLE)
			postCh = make(chan int64)
			go models.InsertPost(mdInfo[TITLE], filepath.Base(filePath), int64(textAmount), postCh)
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
				models.InsertPostAndCategory(postId, v)
			}
		}()
	}

	if labelCh != nil {
		go func() {
			labels := <-labelCh
			log.Println("[INFO] labels: ", labels)

			for _, v := range labels {
				models.InsertPostAndLabel(postId, v)
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

// GetStrLength 返回输入的字符串的字数，汉字和中文标点算 1 个字数，英文和其他字符 2 个算 1 个字数，不足 1 个算 1个
func GetStrLength(str string) float64 {
	var total float64

	reg := regexp.MustCompile("/·|，|。|《|》|‘|’|”|“|；|：|【|】|？|（|）|、/")

	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || reg.Match([]byte(string(r))) {
			total = total + 1
		} else {
			total = total + 0.5
		}
	}

	return math.Ceil(total)
}
