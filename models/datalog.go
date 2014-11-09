package models

import (
	"beezo/common"
	"beezo/events"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type DataLog struct {
	Id               int64
	Time             time.Time
	SensorinoAddress string
	ServiceIndex     int64
	ChannelIndex     int64
	Data             string `orm:"size(1024)"`
}

func (this *DataLog) Create() *common.Error {
	if err := this.Check(); err != nil {
		return err
	}

	if _, err := GetChannel(this.SensorinoAddress, this.ServiceIndex, this.ChannelIndex); err != nil {
		if err.Type == common.SensorinoNotFound {
			events.Publish(
				events.Event{
					Type:             events.MissedSensorinoEvent,
					SensorinoAddress: this.SensorinoAddress,
					ServiceIndex:     this.ServiceIndex,
					ChannelIndex:     this.ChannelIndex,
				})
		}
		return common.NewError(fmt.Sprintf("Unable to find channel %+v", this), common.ChannelNotFound)
	}

	// insert
	o := orm.NewOrm()
	_, err := o.Insert(this)

	return common.ConvertError(err, common.X)
}

func (this *DataLog) Delete() *common.Error {
	return common.NewError("TO BE SUPPLIED LATER", common.X)
}

func (this *DataLog) Check() *common.Error {
	if validAddress.MatchString(this.SensorinoAddress) == false {
		return common.NewError("Invalid address for sensorino", common.X)
	}

	return nil
}
