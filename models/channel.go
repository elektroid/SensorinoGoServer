package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	_ "time"
)

type Channel struct {
	Id               int64
	SensorinoAddress string
	ServiceIndex     int64
	Index            int64
	Type             string `orm:"size(100)"`
	DataType         string `orm:"size(100)"`
}

func GetChannel(address string, serviceIndex int64, index int64) (Channel, error) {
	channel := Channel{}
	o := orm.NewOrm()
	err := o.QueryTable("channel").Filter("SensorinoAddress", address).Filter("ServiceIndex", serviceIndex).Filter("Id", index).One(&channel)
	return channel, err
}

func GetChannels(address string, serviceIndex int64) ([]Channel, error) {
	o := orm.NewOrm()
	var channels []Channel
	_, err := o.QueryTable("channel").Filter("SensorinoAddress", address).Filter("ServiceIndex", serviceIndex).All(&channels)
	return channels, err
}

func (this *Channel) Create() error {
	if err := this.Check(); err != nil {
		return err
	}

	if s, err := GetChannel(this.SensorinoAddress, this.ServiceIndex, this.Index); err != nil {
		return errors.New(fmt.Sprintf("Channel with same index already exists, db id: %s", s.Id))
	}

	// insert
	o := orm.NewOrm()
	_, err := o.Insert(&this)
	return err
}

func (this *Channel) Delete() error {
	return errors.New("TO BE SUPPLIED LATER")
}

func (this *Channel) Update() error {
	return errors.New("Operation not supported (need to code logs update too)")
}

func (this *Channel) Check() error {
	if validAddress.MatchString(this.SensorinoAddress) == false {
		return errors.New("Invalid address for sensorino")
	}
	if this.DataType == "" {
		return errors.New("No datatype")
	}
	if this.Type == "" {
		return errors.New("No type: ro or rw")
	}

	return nil
}