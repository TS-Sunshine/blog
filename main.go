package main

import (
	"blog/controllers"
	_ "blog/models"
	_ "blog/routers"
	"github.com/astaxie/beego"
	"strings"
)

//入口函数
func main() {
	// beego 自带错误处理
	beego.ErrorController(&controllers.ErrorController{})
	initSession()
	initTemplate()
	beego.Run()
}

func initTemplate()  {
	beego.AddFuncMap("equrl", func(path, url string) bool {
		tempPath := strings.Trim(path, "/")
		tempUrl := strings.Trim(url, "/")
		return strings.Compare(tempPath, tempUrl) == 0
	})
}

func initSession()  {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

