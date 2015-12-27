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
	Id       int64 `orm:"pk;auto"`
	Fullname string
	Age      string
	Gender   string
	Mobile   string
	Email    string
	Workyear string
	City     string
	Post     string
	Md5      string    `orm:"unique"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

func (c Candidate) Get() error {
	o := orm.NewOrm()
	err := o.Read(c)
	return err
}

func (c *Candidate) getMD5() {
	s := c.Fullname + c.Age + c.Gender + c.Mobile + c.Email + c.Workyear + c.Post + c.City
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

func (c Candidate) GetCandidates() []orm.ParamsList {
	o := orm.NewOrm()
	lists := []orm.ParamsList{}
	o.QueryTable("candidate").ValuesList(&lists, "Id", "Fullname", "Age", "Gender", "Mobile", "Email", "Workyear", "City", "Post", "md5")
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
