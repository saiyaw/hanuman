package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/models"
)

type AttachmentController struct {
	beego.Controller
}

func (c *AttachmentController) InsertOneAttachment() {
	var attachment models.Attachment
	attachment.Filepath = c.GetString("filepath")
	attachment.Candidateid, _ = c.GetInt64("candidateid")
	attachment.Insert()
	c.Ctx.WriteString(strconv.FormatInt(attachment.Id, 10))

}

func (c *AttachmentController) GetAttachments() {
	var attachment models.Attachment
	attachment.Candidateid, _ = c.GetInt64("candidateid")
	ps := attachment.GetAttachments()
	c.Data["json"] = &ps
	c.ServeJson()
}
