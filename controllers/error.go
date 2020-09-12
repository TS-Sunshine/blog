package controllers

import (
	"blog/syserror"
	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404()  {
	c.Data["content"] = "page not found"
	c.TplName = "error/404.html"
	if c.IsAjax() {
		c.jsonerror(syserror.Error404{})
	}

}

func (c *ErrorController) Error500()  {
	c.TplName = "500.html"
	//强转系统error
	err, ok := c.Data["error"].(error)
	if !ok {
		err = syserror.NewError("位置错位", nil)
	}
	//强转自定义error
	serr, ok := err.(syserror.Error)
	if !ok {
		serr = syserror.NewError(err.Error(), nil)
	}
	if serr.ReasonError() != nil {
		logs.Info(serr.Error(), serr.ReasonError())
	}
}

func (c *ErrorController ) jsonerror(serr syserror.Error)  {
	c.Ctx.Output.Status = 200
	c.Data["json"] = map[string]interface{} {
		"code" : 1002,
		"msg" : "非法访问",
	}
	c.ServeJSON()
}
