package routers

import (
	"hdy/music/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    netEaseNameSpace:=beego.NewNamespace("/netease",
    	beego.NSRouter("/search",&controllers.NetEaseController{},"*:Search"),
		beego.NSRouter("/index",&controllers.NetEaseController{},"*:Index"))
    beego.AddNamespace(netEaseNameSpace)
}
