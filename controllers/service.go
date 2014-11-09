package controllers

import (
	"beezo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type ServiceController struct {
	beego.Controller
}

func (this *ServiceController) Post() {
	var service models.Service
	errj := json.Unmarshal(this.Ctx.Input.RequestBody, &service)
	if errj != nil {
		this.Data["json"] = errj
		this.ServeJson()
	}

	// insert
	o := orm.NewOrm()
	_, err := o.Insert(&service)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = service //map[string]string{"ObjectId": objectId}
	}
	this.ServeJson()
}

func (this *ServiceController) Get() {
	sensorinoAddress := this.Ctx.Input.Params[":sensorinoAddress"]
	serviceIndex, err := strconv.ParseInt(this.Ctx.Input.Params[":sensorinoAddress"], 10, 64)
	if err != nil {
		this.Data["json"] = "malformed url"
		this.ServeJson()
	}

	service, err := models.GetService(sensorinoAddress, serviceIndex)

	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = service //map[string]string{"ObjectId": objectId}
	}
	this.ServeJson()
}

func (this *ServiceController) Put() {
	var service models.Service
	if errj := json.Unmarshal(this.Ctx.Input.RequestBody, &service); errj != nil {
		this.Data["json"] = errj
		this.ServeJson()
		return
	}

	err := service.Update()
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = service //map[string]string{"ObjectId": objectId}
	}
	this.ServeJson()
}

func (this *ServiceController) Delete() {
	sensorinoAddress := this.Ctx.Input.Params[":sensorinoAddress"]
	serviceIndex, err := strconv.ParseInt(this.Ctx.Input.Params[":sensorinoAddress"], 10, 64)
	if err != nil {
		this.Data["json"] = "malformed url"
		this.ServeJson()
	}

	this.Data["json"] = fmt.Sprintf("TBSL not deleted %s, %d", sensorinoAddress, serviceIndex)
	this.ServeJson()
}
