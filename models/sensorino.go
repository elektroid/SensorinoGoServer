package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"regexp"
)

type Sensorino struct {
	Id      int64
	Name    string `orm:"size(100)" json:"Name"`
	Address string `orm:"size(100)"`
}

func GetSensorino(address string) (Sensorino, error) {
	senso := Sensorino{}
	o := orm.NewOrm()
	err := o.QueryTable("sensorino").Filter("Address", address).One(&senso)
	return senso, err
}

func (this *Sensorino) Create() error {
	if err := this.Check(); err != nil {
		return err
	}

	if s, err := GetSensorino(this.Address); err != nil {
		return errors.New(fmt.Sprintf("Sensorino with same address exists %s", s.Name))
	}

	// insert
	o := orm.NewOrm()
	_, err := o.Insert(&this)
	return err
}

func (this *Sensorino) Delete() error {
	return errors.New("TO BE SUPPLIED LATER")
}

func (this *Sensorino) Update() error {

	var err error
	// if it's invalid, don't bother
	if err = this.Check(); err != nil {
		return err
	}

	// if it has no Id field, it was loaded from outside
	if this.Id == 0 {
		var s Sensorino
		if s, err = GetSensorino(this.Address); err != nil {
			return errors.New("Changing address is not supported yet")
		}
		this = &s
	}

	// insert
	o := orm.NewOrm()
	_, err = o.Update(&this)
	return err

}

var validAddress = regexp.MustCompile(`^\d+\.\d+\.\d+.\d$`)

func (this *Sensorino) Check() error {
	if this.Name == "" {
		return errors.New("Invalid name for sensorino")
	}
	if validAddress.MatchString(this.Address) == false {
		return errors.New("Invalid address for sensorino")
	}
	return nil
}
