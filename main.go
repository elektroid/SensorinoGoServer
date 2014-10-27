package main

import (
	"beezo/controllers"
	"beezo/models"
	_ "beezo/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {

	models.InitOrm()
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	o := orm.NewOrm()
	senso := models.Sensorino{Name: "slene", Address: "1.2.3.4"}
	// insert
	id, err := o.Insert(&senso)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	beego.RESTRouter("/sensorino", &controllers.SensorinoController{})
	beego.Run()
}
