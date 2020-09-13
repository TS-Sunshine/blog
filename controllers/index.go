package controllers

type IndexController struct {
	BaseController
}

// @router / [get]
func (this *IndexController) GetIndex() {
	this.TplName = "index.html"
}

// @router /message [get]
func (this *IndexController) GetMessage() {
	this.TplName = "message.html"
}

// @router /about [get]
func (this *IndexController) GetAbout()  {
	this.TplName = "about.html"
}

// @router /user [get]
func (this *IndexController) GetUser()  {
	this.TplName = "user.html"
}