package controllers

import (
	"github.com/astaxie/beego"
	"hdy/music/models"
)

type NetEaseController struct {
	beego.Controller
}

func (c *NetEaseController) Search() {
	q:=c.GetString("q")
	p:=c.GetString("p")
	res:=(&models.NetEase{}).Search(p,q)
	c.Data["json"] = &res
	c.ServeJSON()
}
func (c *NetEaseController) Index() {
	c.TplName=""
}