package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "sensorino.org"
	this.Data["Email"] = "bozo@nobodix.org"
	this.TplNames = "index.tpl"
}
