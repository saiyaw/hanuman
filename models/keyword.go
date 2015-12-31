package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Keyword struct {
	Id      int64     `orm:"pk;auto"`
	Content string    `orm:"unique"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (k *Keyword) Insert() error {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(k)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	} else {
		k.Id = id
	}
	o.Commit()
	return nil
}

func (k Keyword) Get() error {
	o := orm.NewOrm()
	err := o.Read(k)
	return err
}

func (k Keyword) GetKeywordList() []orm.ParamsList {
	o := orm.NewOrm()
	lists := []orm.ParamsList{}
	o.QueryTable("keyword").ValuesList(&lists, "Id", "Content")
	return lists
}
