package models

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Candidate struct {
	Id        int64 `orm:"pk;auto"`
	Fullname  string
	Age       string
	Gender    string
	Mobile    string
	Email     string
	Workyear  string
	City      string
	Post      string
	Company   string
	Education string
	Birthday  time.Time `orm:"null;type(date)"`
	Md5       string    `orm:"unique"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
}

func (c *Candidate) getMD5() {
	s := c.Fullname + c.Age + c.Gender + c.Mobile + c.Email + c.Workyear + c.Post + c.City + c.Company + c.Education
	hasher := md5.New()
	hasher.Write([]byte(s))
	c.Md5 = hex.EncodeToString(hasher.Sum(nil))
}

func (c *Candidate) Insert() error {
	c.getMD5()
	o := orm.NewOrm()
	o.Begin()
	id, err := o.Insert(c)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	} else {
		c.Id = id
	}
	o.Commit()
	return nil
}

func (c *Candidate) Update() error {
	c.Updated = time.Now()
	c.getMD5()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(c, "Fullname", "Age", "Post", "Workyear", "City", "Gender", "Mobile", "Email", "Company", "Education", "Birthday", "Md5", "Updated")
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return err
}

func (c Candidate) GetCandidates() []orm.ParamsList {
	o := orm.NewOrm()
	lists := []orm.ParamsList{}
	o.QueryTable("candidate").ValuesList(&lists, "Id", "Fullname", "Age", "Post", "Workyear", "City", "Company", "Gender", "Mobile", "Email", "Education", "Birthday", "Md5")
	return lists
}

func (c *Candidate) GetCandidateByID() error {
	o := orm.NewOrm()
	err := o.QueryTable("candidate").Filter("Id", c.Id).One(c)
	return err
}

func (c *Candidate) GetCandidateByMD5() error {
	o := orm.NewOrm()
	err := o.QueryTable("candidate").Filter("Md5", c.Md5).One(c)
	return err
}

func (c Candidate) GetCandidateLabels() string {
	var result string
	o := orm.NewOrm()
	var list orm.ParamsList
	o.Raw("SELECT label.content FROM candidate_label, label WHERE candidate_label.labelid = label.id AND candidateid = ?; ", c.Id).ValuesFlat(&list)
	for _, v := range list {
		result = result + v.(string) + ";"
	}

	return result
}

func GetCandidatesByLabels(ids []int) []Candidate {
	var result []Candidate
	o := orm.NewOrm()
	var list orm.ParamsList
	qs := o.QueryTable("candidate_label")
	qs.GroupBy("candidateid").Filter("labelid__in", ids).ValuesFlat(&list, "candidateid")
	for _, v := range list {
		var c Candidate
		c.Id = v.(int64)
		c.GetCandidateByID()
		result = append(result, c)
	}
	return result
}
