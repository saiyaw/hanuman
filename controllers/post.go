package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/models"
)

type PostController struct {
	beego.Controller
}

func (c *PostController) InsertOnePost() {
	var post models.Post
	post.Title = c.GetString("post")
	post.Insert()
	c.Ctx.WriteString(strconv.FormatInt(post.Id, 10))

}

func (c *PostController) GetPostList() {
	var post models.Post
	ps := post.GetPostList()
	c.Data["json"] = &ps
	c.ServeJSON()
}
