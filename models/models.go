package models

import (
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
)

func init() {
	orm.Debug = false
	orm.RegisterDriver("postgres", orm.DR_Postgres)
	connstr := "user=hanuman password=123456 dbname=hanuman sslmode=disable"
	orm.RegisterDataBase("default", "postgres", connstr)

	orm.RegisterModel(new(Candidate), new(Comment), new(Attachment), new(Post))

	createTables()
}

func createTables() error {
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	return err
}
