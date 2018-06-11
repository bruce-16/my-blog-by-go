package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "bpost"
	app.Usage = "Managing blog posts"
	app.Version = "0.0.1"

	var filePath string
	var fileName string
	var host string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "upload,up",
			Usage:       "Upload an article by that `file path`",
			Destination: &filePath,
		},
		cli.StringFlag{
			Name:        "remove,rm",
			Usage:       "Remove the article by that `file name`",
			Destination: &fileName,
		},
		cli.StringFlag{
			Name:        "config,cf",
			Usage:       "Config http request `host`",
			Destination: &host,
		},
	}
	var url = "http://localhost:8888"
	isexists := exists("./url.config")
	if isexists {
		file, rdRrr := ioutil.ReadFile("./url.config")
		if rdRrr != nil {
			log.Fatalln("[ERROR] ", rdRrr)
		}
		url = string(file)
	}

	app.Action = func(c *cli.Context) error {
		if filePath != "" {
			upload(filePath, url)
		}
		if fileName != "" {
			remove(fileName, url)
		}
		if host != "" {
			config(host)
		}
		return nil
	}
	err := app.Run(os.Args)
	checkErr(err)
}

func upload(filePath, url string) {
	_, err := os.Stat(filePath) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) == false {
			log.Fatalf("%s 路径不存在", filePath)
		}
	}
	// 创建表单文件
	// CreateFormFile 用来创建表单，第一个参数是字段名，第二个参数是文件名
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	fileName := filepath.Base(filePath)
	formFile, err := writer.CreateFormFile("upload", fileName)
	if err != nil {
		log.Fatalf("Create form file failed: %s\n", err)
	}

	// 从文件读取数据，写入表单
	srcFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Open source file failed: %s\n", err)
	}
	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		log.Fatalf("Write to form file falied: %s\n", err)
	}

	// 发送表单
	contentType := writer.FormDataContentType()
	writer.Close() // 发送之前必须调用Close()以写入结尾行
	var res *http.Response
	res, err = http.Post(url+"/upload", contentType, buf)
	if err != nil {
		log.Fatalf("Post failed: %s\n", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var dataMap map[string]interface{}
	json.Unmarshal(body, &dataMap)
	if dataMap["status"] != 0 {
		log.Fatalf("Post failed: %s\n", dataMap["msg"])
	}
	log.Println("[SUCCESS] Upload file is successfully. ", filePath)
}

func remove(fileName, url string) {
	res, err := http.Get(url + "/remove?name=" + fileName)
	if err != nil {
		log.Fatalf("Remove failed: %s\n", err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var dataMap map[string]interface{}
	json.Unmarshal(body, &dataMap)
	if dataMap["status"] != 0 {
		log.Fatalf("Post failed: %s\n", dataMap["msg"])
	}
	fmt.Println("[SUCCESS] Remove file is successfully. ", fileName)
}

func config(host string) {
	var file *os.File
	var err error
	file, err = os.OpenFile("./url.config", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	defer file.Close()

	if err != nil {
		log.Fatalln("[ERROR] ", err)
	}

	_, err = file.WriteString(host)
	if err != nil {
		log.Fatalln("[ERROR] ", err)
	} else {
		log.Fatalln("[SUCCESSFUL] 配置成功！")
	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln("ERROR:", e)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
