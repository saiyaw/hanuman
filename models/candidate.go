package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Candidate struct {
	Id       int64 `orm:"pk;auto"`
	Fullname string
	Age      int64
	Gender   string
	Mobile   string
	Email    string `orm:"unique"`
	Workyear int64
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

func (c Candidate) Get() error {
	o := orm.NewOrm()
	err := o.Read(c)
	return err
}

func (c *Candidate) Insert() error {
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
	o.QueryTable("candidate").ValuesList(&lists, "Id", "Fullname", "Age", "Gender", "Mobile", "Email", "Workyear")
	return lists
}
