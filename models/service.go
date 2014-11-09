package models

import (
	"beezo/common"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type Service struct {
	Id               int64
	SensorinoAddress string
	Index            int64
	Name             string
}

func GetService(address string, index int64) (Service, *common.Error) {
	service := Service{}
	o := orm.NewOrm()
	err := o.QueryTable("service").Filter("SensorinoAddress", address).Filter("Index", index).One(&service)
	return service, common.ConvertError(err, common.X)
}

func GetServices(address string) ([]*Service, *common.Error) {
	o := orm.NewOrm()
	var services []*Service
	_, err := o.QueryTable("service").Filter("SensorinoAddress", address).All(&services)
	return services, common.ConvertError(err, common.X)
}

func (this *Service) Create() *common.Error {
	if err := this.Check(); err != nil {
		return err
	}

	// Sensorino exists ?
	if _, err := GetSensorino(this.SensorinoAddress); err != nil {
		return common.NewError(fmt.Sprintf("Unable to find Sensorino to attach service to %s", this.SensorinoAddress), common.X)
	}

	// Which index for this service ? Let's check existing ones
	// TODO we could check for name collision
	services, errServices := GetServices(this.SensorinoAddress)
	if errServices == nil {
		this.Index = int64(len(services))
	}

	// insert
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return common.ConvertError(err, common.X)
}

func (this *Service) Delete() *common.Error {
	return common.NewError("TO BE SUPPLIED LATER", common.X)
}

func (this *Service) Update() *common.Error {

	var err error
	// if it's invalid, don't bother
	if err = this.Check(); err != nil {
		return common.ConvertError(err, common.X)
	}

	// if it has no Id field, it was loaded from outside
	if this.Id == 0 {
		var s Service
		if s, err = GetService(this.SensorinoAddress, this.Index); err != nil {
			return common.NewError("Changing address or index not supported yet", common.X)
		}
		this = &s
	}

	o := orm.NewOrm()
	_, err = o.Update(this)
	return common.ConvertError(err, common.X)

}

func (this *Service) Check() *common.Error {
	if this.SensorinoAddress == "" {
		return common.NewError("Invalid service: no sensorino id", common.X)
	}
	_, err := GetSensorino(this.SensorinoAddress)
	if err != nil {
		return common.NewError("Failed to load sensorino service is attached to", common.SensorinoNotFound)
	}
	return nil

}
