package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Attachment struct {
	Id          int64 `orm:"pk;auto"`
	Candidateid int64
	Filepath    string
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func (a *Attachment) Insert() error {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(a)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	} else {
		a.Id = id
	}
	o.Commit()
	return nil
}

func (a Attachment) GetAttachments() []orm.ParamsList {
	o := orm.NewOrm()
	lists := []orm.ParamsList{}
	o.QueryTable("attachment").Filter("Candidateid", a.Candidateid).OrderBy("-Created").ValuesList(&lists, "Id", "Candidateid", "Filepath", "Created", "Updated")
	return lists
}
