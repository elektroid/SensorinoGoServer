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

	beego.RESTRouter("/sensorino", &controllers.SensorinoController{})
	beego.Run()
}
