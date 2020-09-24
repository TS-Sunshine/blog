package controllers

import (
	"blog/models"
	"blog/syserror"
	"errors"
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type H map[string]interface{}

type NestPrepare interface {
	NextPrepare()
}

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
	this.Data["bLogin"] = this.bIsLogin
	if a, ok := this.AppController.(NestPrepare); ok {
		a.NextPrepare()
	}
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

func (c *BaseController) MustLogin()  {
	if !c.bIsLogin {
		c.Abort500(syserror.NoUserError{})
	}
}

func (this *BaseController) JsonOK(msg ,action string)  {
	this.Data["json"] = H {
		"code" : 0,
		"msg" : msg,
		"action" : action,
	}
	this.ServeJSON()
}

func (this *BaseController) UUID() string {
	u := uuid.NewV4()
	if len(u.String()) <= 0 {
		this.Abort500(errors.New("系统错误"))
	}
	return u.String()
}