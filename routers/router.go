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
    //新版路由
    //搜索
	sou:=beego.NewNamespace("/sou",
		beego.NSRouter("/search",&controllers.SouController{},"*:Search"),
		beego.NSRouter("/desc",&controllers.SouController{},"*:GetDesc"))
	beego.AddNamespace(sou)
}
