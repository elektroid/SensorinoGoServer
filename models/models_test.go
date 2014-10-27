package models_test

import (
	"beezo/models"
	_ "fmt"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

func TestCreation(t *testing.T) {
	senso := models.Sensorino{Name: "Little one", Address: "1.2.3.4"}
	if err := senso.Create(); err != nil {
		t.Error("Failed to create sensorino")
		t.FailNow(err)
	}

	s2, err2:=GetSensorino(senso.Address)
	if err2 != nil {
		t.Error("Failed to load sensorino back after creation")
		t.FailNow(err)	
	}

	s2.Name="Old one"
	err=s2.Update()
	if err!=nil{
		t.Error("Failed to update sensorino")
		t.FailNow(err)		
	}


	service.
}
