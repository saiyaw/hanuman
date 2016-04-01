package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("postgres", orm.DRPostgres)
	connstr := "user=hanuman password=123456 sslmode=disable dbname=" + beego.AppConfig.String("dbname")
	orm.RegisterDataBase("default", "postgres", connstr)

	orm.RegisterModel(new(Candidate), new(Comment), new(Attachment), new(Post), new(Label), new(CandidateLabel))

	createTables()
}

func createTables() error {
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	return err
}
