package controllers

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/models"
)

type CandidateLabelController struct {
	beego.Controller
}

func (c *CandidateLabelController) InsertCandidateLabel() {

	candidateid, _ := c.GetInt64("candidateid")

	labels := c.GetString("labels")

	array := strings.Split(labels, "|")
	var cls []models.CandidateLabel
	for _, v := range array {
		if len(v) > 0 {
			var cl models.CandidateLabel
			cl.Candidateid = candidateid
			cl.Labelid, _ = strconv.ParseInt(v, 10, 64)
			cls = append(cls, cl)
		}
	}
	models.InsertCandidateLabels(cls)

	c.Ctx.WriteString("ok")
}

func (c *CandidateLabelController) DeleteCandidateLabel() {
	candidateid, _ := c.GetInt64("candidateid")
	var cl models.CandidateLabel
	cl.Candidateid = candidateid
	cl.DeleteCandidateLabels()

	c.Ctx.WriteString("ok")
}

func (c *CandidateLabelController) GetCandidateLabelList() {
	var cl models.CandidateLabel
	cl.Candidateid, _ = c.GetInt64("candidateid")
	ps := cl.GetCandidateLabelList()
	c.Data["json"] = &ps
	c.ServeJson()
}
