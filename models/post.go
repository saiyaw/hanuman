package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id       int64     `orm:"pk;auto"`
	Position string    `orm:"unique"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

func (p *Post) Insert() error {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(p)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	} else {
		p.Id = id
	}
	o.Commit()
	return nil
}

func (p Post) Get() error {
	o := orm.NewOrm()
	err := o.Read(p)
	return err
}

func (p Post) GetPostList() []orm.ParamsList {
	o := orm.NewOrm()
	lists := []orm.ParamsList{}
	o.QueryTable("post").ValuesList(&lists, "Id", "Position", "Created", "Updated")
	return lists
}
