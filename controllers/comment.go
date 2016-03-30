package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/models"
)

type CommentController struct {
	beego.Controller
}

func (c *CommentController) InsertOneComment() {
	var comment models.Comment
	comment.Content = c.GetString("content")
	comment.Candidateid, _ = c.GetInt64("candidateid")
	comment.Insert()
	c.Ctx.WriteString(strconv.FormatInt(comment.Id, 10))

}

func (c *CommentController) GetComments() {

	var comment models.Comment
	comment.Candidateid, _ = c.GetInt64("candidateid")
	ps := comment.GetComments()
	c.Data["json"] = &ps
	c.ServeJSON()
}
