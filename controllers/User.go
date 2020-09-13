package controllers

import (
	"blog/models"
	"blog/syserror"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (this *UserController) Login()  {
	//email
	email := this.GetMustString("email", "邮箱不能为空")
	pwd := this.GetMustString("password", "密码不能为空")
	user, err := models.QueryUserByEmailAndPwd(email, pwd)
	if err != nil {
		this.Abort500(syserror.NewError("登录失败", nil))
	}
 	this.SetSession(SESSION_USER_KEY, user)
	this.Data["json"] = map[string]interface{} {
		"code" : 0,
		"action" : "/",
	}
	this.ServeJSON()
}

// @router /logout [get]
func (this *UserController) Logout() {
	this.MustLogin()
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/", 302)
}