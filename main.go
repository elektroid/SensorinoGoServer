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

	// For REST api we need 2 beego.Router by url
	//
	//first for listing (GET without params)
	//second for targeted operations

	beego.Router("/sensorino", &controllers.SensorinoController{})
	beego.Router("/sensorino/:sensorinoAddress", &controllers.SensorinoController{})

	beego.Router("/sensorino/:id:int64/service", &controllers.ServiceController{})
	beego.Router("/sensorino/:sensorinoAddress/service/:serviceIndex", &controllers.ServiceController{})

	/*
		beego.Router("/sensorino/:sensorinoAddress/service/:serviceIndex/channel", &controllers.ChannelController{})
		beego.Router("/sensorino/:sensorinoAddress/service/:serviceIndex/channel/:channelIndex", &controllers.ChannelController{})
	*/

	beego.Run()
}
