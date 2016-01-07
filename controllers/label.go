package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/models"
)

type LabelController struct {
	beego.Controller
}

func (c *LabelController) InsertOneLabel() {
	var l models.Label
	l.Content = c.GetString("label")
	l.Insert()
	c.Ctx.WriteString(strconv.FormatInt(l.Id, 10))
}

func (c *LabelController) DeleteOneLabel() {
	var l models.Label
	l.Id, _ = c.GetInt64("labelid")
	err := l.SetDeleted()
	if err != nil {
		c.Ctx.WriteString(err.Error())
	} else {
		c.Ctx.WriteString("ok")
	}
}

func (c *LabelController) GetLabelList() {
	var l models.Label
	ps := l.GetLabelList()
	c.Data["json"] = &ps
	c.ServeJson()
}

func (c *LabelController) GetLabel() {
	var l models.Label
	l.Id, _ = c.GetInt64("labelid")
	l.Get()
	c.Ctx.WriteString(l.Content)
}

func (c *LabelController) CleanUselessLabels() {
	var l models.Label
	lst := l.GetSetDeletedLabelList()
	for _, v := range lst {
		var candidatelabel models.CandidateLabel
		candidatelabel.Labelid = v[0].(int64)
		if candidatelabel.IsExistingInCandidateLabel() {
			// do nothing
		} else {
			var label models.Label
			label.Id = v[0].(int64)
			label.Delete()
		}
	}
	c.Ctx.WriteString("ok")
}
