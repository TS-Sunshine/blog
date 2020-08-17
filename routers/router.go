package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	/*首页 index*/
	//beego.Router("/", &controllers.ViewController{})
	beego.Include(&controllers.IndexController{})
}
