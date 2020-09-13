package controllers

import (
	"blog/models"
	"errors"
	"github.com/astaxie/beego"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type BaseController struct {
	beego.Controller
	User models.User
	bIsLogin bool
}

func (this *BaseController) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	if ok {
		this.User = u
		this.bIsLogin = true
		this.Data["User"] = this.User
	}
	//this.Data["islogin"] = this.bIsLogin
}

func (this *BaseController) Abort500(err error)  {
	this.Data["error"] = err
	this.Abort("500")
}

func (this *BaseController) GetMustString(key, msg string) string {
	value := this.GetString(key)
	if len(value) == 0 {
		this.Abort500(errors.New(msg))
	}
	return value
}