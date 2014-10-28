package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func InitOrm() {

	// register model
	orm.RegisterModel(new(Sensorino))
	orm.RegisterModel(new(Service))
	orm.RegisterModel(new(Channel))
	orm.RegisterModel(new(DataLog))

	// set default database
	orm.RegisterDataBase("default", "sqlite3", "./test.db", 30)
}
