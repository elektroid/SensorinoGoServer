package models

import (
	_ "fmt"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type DataLog struct {
	Id        int64
	Time      time.Time
	ChannelId int
	Data      string `orm:"size(1024)"`
}
