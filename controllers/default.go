package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
	c.Layout = "layout.tpl"
}

type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {
	c.TplName = "list.tpl"
	c.Layout = "layout.tpl"
}

type NewController struct {
	beego.Controller
}

func (c *NewController) Get() {
	c.TplName = "basic.tpl"
	c.Layout = "layout.tpl"

}

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.TplName = "admin.tpl"
	c.Layout = "layout.tpl"
}
