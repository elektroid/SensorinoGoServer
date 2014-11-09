// Beego (http://beego.me/)
// @description beego is an open-source, high-performance web framework for the Go programming language.
// @link        http://github.com/astaxie/beego for the canonical source repository
// @license     http://github.com/astaxie/beego/blob/master/LICENSE
// @authors     astaxie

package controllers

import (
	"beezo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type SensorinoController struct {
	beego.Controller
}

func (this *SensorinoController) Post() {
	var senso models.Sensorino
	errj := json.Unmarshal(this.Ctx.Input.RequestBody, &senso)
	if errj != nil {
		this.Data["json"] = errj
		this.ServeJson()
	}

	err := senso.Create()
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = senso //map[string]string{"ObjectId": objectId}
	}
	this.ServeJson()
}

func (this *SensorinoController) Get() {
	sensorinoAddress := this.Ctx.Input.Params[":sensorinoAddress"]

	if sensorinoAddress == "" {
		sensorinos, err := models.GetSensorinos()
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = sensorinos
		}
	} else {

		s, err := models.GetSensorino(sensorinoAddress)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = s
		}

	}
	this.ServeJson()
}

func (this *SensorinoController) Put() {
	objectId := this.Ctx.Input.Params[":objectId"]

	var senso models.Sensorino
	json.Unmarshal(this.Ctx.Input.RequestBody, &senso)
	o := orm.NewOrm()
	num, err := o.Update(&senso)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = fmt.Sprintf("update success : %d", objectId)
	}
	this.ServeJson()
}

func (this *SensorinoController) Delete() {
	pobjectId := this.Ctx.Input.Params[":objectId"]
	objectId, convErr := strconv.ParseInt(pobjectId, 10, 32)
	if convErr != nil {
		this.Data["json"] = convErr
	} else {

		// delete
		s := models.Sensorino{Id: objectId}
		o := orm.NewOrm()
		_, err := o.Delete(&s)
		if err != nil {
			this.Data["json"] = "delete success!"
		} else {
			this.Data["json"] = err
		}
	}

	this.ServeJson()
}
