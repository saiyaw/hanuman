package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplNames = "index.tpl"
	c.Layout = "layout.tpl"
}

type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {
	c.TplNames = "list.tpl"
	c.Layout = "layout.tpl"
}

type NewController struct {
	beego.Controller
}

func (c *NewController) Get() {
	c.TplNames = "basic.tpl"
	c.Layout = "layout.tpl"

}

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.TplNames = "admin.tpl"
	c.Layout = "layout.tpl"
}
