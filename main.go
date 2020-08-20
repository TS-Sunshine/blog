package main

import (
	_ "blog/models"
	_ "blog/routers"
	"github.com/astaxie/beego"
	"strings"
)

func main() {
	initTemplate()
	beego.Run()
}

func initTemplate()  {
	beego.AddFuncMap("equrl", func(path, url string) bool {
		tempPath := strings.Trim(path, "/")
		tempUrl := strings.Trim(url, "/")

		return strings.Compare(tempPath, tempUrl) == 0
	});
}

