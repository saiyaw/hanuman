package controllers

import (
	"github.com/astaxie/beego"
)

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
	c.TplNames = "new.tpl"
	c.Layout = "layout.tpl"
}

type ViewController struct {
	beego.Controller
}

func (c *ViewController) Get() {
	c.TplNames = "view.tpl"
	c.Layout = "layout.tpl"
}
