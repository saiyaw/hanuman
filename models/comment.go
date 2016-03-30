package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Comment struct {
	Id          int64 `orm:"pk;auto"`
	Candidateid int64
	Content     string    `orm:"type(text)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func (c *Comment) Insert() error {
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

func (c Comment) GetComments() []orm.ParamsList {
	o := orm.NewOrm()
	lists := []orm.ParamsList{}
	o.QueryTable("comment").Filter("Candidateid", c.Candidateid).OrderBy("-Created").ValuesList(&lists, "Id", "Candidateid", "Content", "Created", "Updated")
	return lists
}
