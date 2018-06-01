package controllers

import (
	"github.com/astaxie/beego"
	"hdy/music/models"
)
type SouController struct {
	beego.Controller
}

var searcher models.Searcher

func (c *SouController)Search()  {
	which:= c.GetString("type")
	q:= c.GetString("q")
	p:= c.GetString("p")
	if q==""{
		data:="fail_no_key"
		c.Data["json"] = &data
	}else {
		switch which {
		case "":
			fallthrough
		case "netease":
			searcher=&models.NetEase{}
		case "kuwo":
			searcher=&models.KuWo{}
		}
		res:=searcher.Search(q,p)
		c.Data["json"] = &res
	}
	c.ServeJSON()
}
func (c *SouController)GetDesc()  {
	which:= c.GetString("type")
	id:= c.GetString("id")
	if id==""{
		data:="fail_no_id"
		c.Data["json"] = &data
	}else {
		switch which {
		case "":
			fallthrough
		case "netease":
			searcher=&models.NetEase{}
		case "kuwo":
			searcher=&models.KuWo{}
		}
		res:=searcher.GetDesc(id)
		c.Data["json"] = &res
	}
	c.ServeJSON()
}
