package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type CandidateLabel struct {
	Id          int64 `orm:"pk;auto"`
	Candidateid int64
	Labelid     int64
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func (cl *CandidateLabel) Insert() error {
	o := orm.NewOrm()
	o.Begin()

	id, err := o.Insert(cl)
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	} else {
		cl.Id = id
	}
	o.Commit()
	return nil
}

func (cl *CandidateLabel) Get() error {
	o := orm.NewOrm()
	err := o.Read(cl)
	return err
}

func InsertCandidateLabels(Labels []CandidateLabel) {
	o := orm.NewOrm()

	qs := o.QueryTable("candidate_label")
	i, _ := qs.PrepareInsert()
	for _, cl := range Labels {
		_, err := i.Insert(&cl)
		if err != nil {
			log.Println("insert patient answer failed :" + err.Error())
		}
	}
	i.Close()
}

func (cl CandidateLabel) IsExistingInCandidateLabel() bool {
	o := orm.NewOrm()
	err := o.Read(&cl, "Labelid")
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func (cl CandidateLabel) DeleteCandidateLabels() error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.QueryTable("candidate_label").Filter("Candidateid", cl.Candidateid).Delete()
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

func (cl CandidateLabel) DeleteCandidateLabelsByLabel() error {
	o := orm.NewOrm()
	o.Begin()
	_, err := o.QueryTable("candidate_label").Filter("Labelid", cl.Labelid).Delete()
	if err != nil {
		log.Println(err.Error())
		o.Rollback()
		return err
	}
	o.Commit()
	return nil
}

func (cl CandidateLabel) GetCandidateLabelList() []orm.ParamsList {
	o := orm.NewOrm()
	lists := []orm.ParamsList{}
	o.QueryTable("candidate_label").Filter("Candidateid", cl.Candidateid).ValuesList(&lists, "Id", "Labelid")
	return lists
}
