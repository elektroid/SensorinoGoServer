package models

import (
	"errors"
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

func GetService(address string, index int64) (Service, error) {
	service := Service{}
	o := orm.NewOrm()
	err := o.QueryTable("service").Filter("SensorinoAddress", address).Filter("Index", index).One(&service)
	return service, err
}

func GetServices(address string) ([]*Service, error) {
	o := orm.NewOrm()
	var services []*Service
	_, err := o.QueryTable("service").Filter("SensorinoAddress", address).All(&services)
	return services, err
}

func (this *Service) Create() error {
	if err := this.Check(); err != nil {
		return err
	}

	// Sensorino exists ?
	if _, err := GetSensorino(this.SensorinoAddress); err != nil {
		return errors.New(fmt.Sprintf("Unable to find Sensorino to attach service to %s", this.SensorinoAddress))
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
	return err
}

func (this *Service) Delete() error {
	return errors.New("TO BE SUPPLIED LATER")
}

func (this *Service) Update() error {

	var err error
	// if it's invalid, don't bother
	if err = this.Check(); err != nil {
		return err
	}

	// if it has no Id field, it was loaded from outside
	if this.Id == 0 {
		var s Service
		if s, err = GetService(this.SensorinoAddress, this.Index); err != nil {
			return errors.New("Changing address or index not supported yet")
		}
		this = &s
	}

	o := orm.NewOrm()
	_, err = o.Update(this)
	return err

}

func (this *Service) Check() error {
	if this.SensorinoAddress == "" {
		return errors.New("Invalid service: no sensorino id")
	}
	_, err := GetSensorino(this.SensorinoAddress)
	if err != nil {
		return errors.New("Failed to load sensorino service is attached to")
	}
	return nil

}
