# 使用go搭建一个简易的博客系统

* 后端使用Go语言开发，使用Web框架gin和ORM框架xorm。
* 前端使用React+Metrail-Ui+React-Router4。
* 命令行工具使用`github.com/urfave/cli`开发。

## 目前功能

* 命令行工具能对文章进行上传、删除、和进行请求`Host`进行配置。
* 服务端提供了对于文章的基本增删改查接口，使用`sqlite`数据来存储文章相应的标题、标签和分类等。
* 前端能简单展示文章列表和文章内容，根据标签或者分类展示相关的文章。

## 截图

![展示图1](http://p9uc2ui6z.bkt.clouddn.com/github/my-blog-by-go/1.png)

![展示图2](http://p9uc2ui6z.bkt.clouddn.com/github/my-blog-by-go/2.png)

![展示图3](http://p9uc2ui6z.bkt.clouddn.com/github/my-blog-by-go/3.png)

## 启动项目

* 服务端
```shell
go get -u "https://github.com/zachrey/my-blog-by-go"

cd $GOPATH/src/github.com/zachrey/my-blog-by-go/

go run main.go
```
端口默认开启的是本地的`8888`。

* 前端
```shell
cd $GOPATH/src/github.com/zachrey/my-blog-by-go/front_web

yarn && yarn start
```

* 命令行工具
```shell
cd $GOPATH/src/github.com/zachrey/my-blog-by-go/cmd

go run main.go --help
```

这里，前后端和命令行工具都没有进行编译，直接在开发环境中演示。


> 如果go 项目中有些依赖包下载不下来，建议翻墙或者去github找相应的库，然后将它clone到你的src/github文件夹相应路劲下。

## 项目结构
.
├── cmd              // 命令行工具
├── controllers      // 控制器
├── database         // 数据库连接，配置等
├── front_web        // react前端内容
├── models           // 数据库相关表的操作模板
├── posts            // 存放上传的文章
├── routers          // 路由
├── static           // 静态文件资源，前端打包后就可以打包放到这里面
├── gin.log          // web运行日志
├── main.go          // 程序入口
├── README.md
└── vendor

欢迎讨论和star