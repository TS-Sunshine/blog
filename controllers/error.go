package controllers

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404()  {
	c.Data["content"] = "page not found"
	c.TplName = "404.html"
	if c.IsAjax() {
		c.Ctx.Output.Status = 200
		c.Data["json"] = map[string]interface{} {
			"code" : 1002,
			
			"msg" : "非法访问",
		}
	}
	c.ServeJSON()
}

func (c *ErrorController) Error500()  {
	c.TplName = "500.html"
	err, ok := c.Data["error"].(error)
	if !ok {

	}
}
