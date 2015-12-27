package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/models"
)

type CandidateController struct {
	beego.Controller
}

func (c *CandidateController) Get() {
	c.TplNames = "candidate.tpl"
	c.Layout = "layout.tpl"
}

func (c *CandidateController) InsertOneCandidate() {
	var candidate models.Candidate
	candidate.Fullname = c.GetString("fullname")

	candidate.Age = c.GetString("age")

	candidate.Gender = c.GetString("gender")

	candidate.Mobile = c.GetString("mobile")

	candidate.Email = c.GetString("email")

	candidate.Workyear = c.GetString("workyear")

	candidate.Post = c.GetString("post")

	candidate.City = c.GetString("city")

	candidate.Insert()
	c.Ctx.WriteString(strconv.FormatInt(candidate.Id, 10))

}

func (c *CandidateController) GetCandidates() {

	var candidate models.Candidate
	ps := candidate.GetCandidates()
	c.Data["json"] = &ps
	c.ServeJson()
}

func (c *CandidateController) GetCandidateByMD5() {
	md5 := c.GetString("md5")
	var candidate models.Candidate
	candidate.Md5 = md5
	candidate.GetCandidateByMD5()
	c.Data["json"] = &candidate
	c.ServeJson()
}

func (c *CandidateController) GetCandidateByID() {

	var candidate models.Candidate
	candidate.Id, _ = c.GetInt64("id")
	candidate.GetCandidateByID()
	c.Data["json"] = &candidate
	c.ServeJson()
}
