package models

import (
	"beezo/common"
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

func GetChannel(address string, serviceIndex int64, index int64) (Channel, *common.Error) {
	channel := Channel{}
	o := orm.NewOrm()
	err := o.QueryTable("channel").Filter("SensorinoAddress", address).Filter("ServiceIndex", serviceIndex).Filter("Id", index).One(&channel)
	return channel, common.ConvertError(err, common.X)
}

func GetChannels(address string, serviceIndex int64) ([]Channel, *common.Error) {
	o := orm.NewOrm()
	var channels []Channel
	_, err := o.QueryTable("channel").Filter("SensorinoAddress", address).Filter("ServiceIndex", serviceIndex).All(&channels)
	return channels, common.ConvertError(err, common.X)
}

func (this *Channel) Create() *common.Error {
	if err := this.Check(); err != nil {
		return common.ConvertError(err, common.X)
	}

	// we're supposed to be attached to a service (and a sensorino, but we're not responsible for the whole chain)
	if _, err := GetService(this.SensorinoAddress, this.ServiceIndex); err != nil {

		return common.NewError(
			fmt.Sprintf("Unable to find Service to attach channel to %s", this.SensorinoAddress),
			common.ServiceNotFound,
		)
	}

	channels, errChans := GetChannels(this.SensorinoAddress, this.ServiceIndex)
	if errChans == nil {
		this.Index = int64(len(channels))
	}

	// insert
	o := orm.NewOrm()
	_, err := o.Insert(this)
	return common.ConvertError(err, common.X)
}

func (this *Channel) Delete() *common.Error {
	return common.NewError("TO BE SUPPLIED LATER", common.X)
}

func (this *Channel) Update() *common.Error {
	return common.NewError("Operation not supported (need to code logs update too)", common.X)
}

func (this *Channel) Check() *common.Error {
	if validAddress.MatchString(this.SensorinoAddress) == false {
		return common.NewError("Invalid address for sensorino", common.X)
	}
	if this.DataType == "" {
		return common.NewError("No datatype", common.X)
	}
	if this.Type == "" {
		return common.NewError("No type: ro or rw", common.X)
	}

	return nil
}
