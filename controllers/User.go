package controllers

import (
	"blog/models"
	"blog/syserror"
	"errors"
	"strings"
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
	this.JsonOK("登录成功", "/")
}

// @router /reg [post]
func (this *UserController) Reg()  {
	name := this.GetMustString("name", "昵称不能为空")
	email := this.GetMustString("email", "邮箱不能为空")
	pwd := this.GetMustString("password", "密码不能为空")
	password := this.GetMustString("password2","密码不能为空")

	if strings.Compare(pwd, password) != 0 {
		this.Abort500(errors.New("两次输入密码不一致"))
	}

	if u, err := models.QueryUserByName(name); err == nil && u.ID > 0 {
		this.Abort500(errors.New("用户昵称已存在"))
	}

	if u, err := models.QueryUserByEmail(email); err == nil && u.ID > 0 {
		this.Abort500(errors.New("邮箱已存在"))
	}

	if err :=  models.SaveUser(&models.User{
		Name: name,
		Email: email,
		Pwd: password,
		Avatar: "/static/images/info-img.png",
		Role: 1,
	}); err != nil{
		this.Abort500(errors.New("注册失败"))
	}

}

// @router /logout [get]
func (this *UserController) Logout() {
	this.MustLogin()
	this.DelSession(SESSION_USER_KEY)
	this.Redirect("/", 302)
}