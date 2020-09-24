package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.UserController{})
	beego.AddNamespace(beego.NewNamespace("/note", beego.NSInclude(&controllers.NoteController{})))
}
