package controllers

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/models"
)

type AttachmentController struct {
	beego.Controller
}

func (c *AttachmentController) InsertOneAttachment() {

	workpath, _ := os.Getwd()
	beego.Debug(workpath)
	f, h, err := c.GetFile("attachment_upload")
	f.Close()
	if err != nil {
		beego.Error(err)
	} else {
		c.SaveToFile("attachment_upload", workpath+"/static/file/"+h.Filename)
	}

	var attachment models.Attachment
	attachment.Filepath = h.Filename
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
