package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetIndex",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["blog/controllers:IndexController"] = append(beego.GlobalControllerRouter["blog/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetMessage",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
