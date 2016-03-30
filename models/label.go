package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Label struct {
	Id        int64  `orm:"pk;auto"`
	Content   string `orm:"unique"`
	Isdeleted bool
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
}

func (l *Label) Insert() error {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(l)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	} else {
		l.Id = id
	}
	o.Commit()
	return nil
}

func (l Label) SetDeleted() error {
	l.Isdeleted = true
	l.Updated = time.Now()
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Update(&l, "Isdeleted", "Updated")
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

func (l Label) Delete() error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Delete(&l)
	if err != nil {
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

func (l *Label) Get() error {
	o := orm.NewOrm()
	err := o.Read(l)
	return err
}

func (l Label) GetLabelList() []orm.ParamsList {
	o := orm.NewOrm()
	lists := []orm.ParamsList{}
	o.QueryTable("label").ValuesList(&lists, "Id", "Content", "Isdeleted")
	return lists
}

func (l Label) GetSetDeletedLabelList() []orm.ParamsList {
	o := orm.NewOrm()
	lists := []orm.ParamsList{}
	o.QueryTable("label").Filter("Isdeleted", true).ValuesList(&lists, "Id")
	return lists
}
