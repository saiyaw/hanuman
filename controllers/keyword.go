package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/models"
)

type KeywordController struct {
	beego.Controller
}

func (c *KeywordController) InsertOneKeyword() {
	var keyword models.Keyword
	keyword.Content = c.GetString("label")
	keyword.Insert()
	c.Ctx.WriteString(strconv.FormatInt(keyword.Id, 10))
}

func (c *KeywordController) DeleteOneKeyword() {
	var keyword models.Keyword
	keyword.Id, _ = c.GetInt64("id")
	err := keyword.Delete()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	} else {
		c.Ctx.WriteString("ok")
	}

}

func (c *KeywordController) Getkeyword() {
	var keyword models.Keyword
	keyword.Content = c.GetString("label")
	keyword.Get()
	c.Data["json"] = &keyword
	c.ServeJson()
}

func (c *KeywordController) GetKeywordList() {
	var keyword models.Keyword
	ps := keyword.GetKeywordList()
	c.Data["json"] = &ps
	c.ServeJson()
}
