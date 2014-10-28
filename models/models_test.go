package models_test

import (
	"beezo/models"
	"fmt"
	"github.com/astaxie/beego/orm"

	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func setup() {
	models.InitOrm()
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	//o.Using("default")

}

func TestCreation(t *testing.T) {

	setup()

	senso := models.Sensorino{Name: "Little one", Address: "1.2.3.4"}
	if err := senso.Create(); err != nil {
		t.Error("Failed to create sensorino")
		t.Fatal(err)
	}

	s2, err2 := models.GetSensorino(senso.Address)
	if err2 != nil {
		t.Error("Failed to load sensorino back after creation")
		t.Fatal(err2)
	}

	s2.Name = "Old one"
	err := s2.Update()
	if err != nil {
		t.Error("Failed to update sensorino")
		t.Fatal(err)
	}

	// add some services
	service := models.Service{SensorinoAddress: senso.Address, Name: "Test Service"}
	if err := service.Create(); err != nil {
		t.Error("Failed to create service")
		t.Fatal(err)
	}

	service2 := models.Service{SensorinoAddress: senso.Address, Name: "Test Service"}
	if err := service2.Create(); err != nil {
		t.Error("Failed to create service")
		t.Fatal(err)
	}

	services, errServices := models.GetServices(senso.Address)
	if errServices != nil {
		t.Error("Failed to load back services")
		t.Fatal(errServices)
	}

	if services[0].Index+1 != services[1].Index {
		t.Fatal("index incr not working for services")
	}

	// add some chans
	channel1 := models.Channel{SensorinoAddress: senso.Address, ServiceIndex: service.Index, DataType: "TEMP", Type: "RO"}
	if err := channel1.Create(); err != nil {
		t.Error("Failed to create channel")
		t.Fatal(err)
	}

	channel2 := models.Channel{SensorinoAddress: senso.Address, ServiceIndex: service.Index, DataType: "COUNT", Type: "RO"}
	if err := channel2.Create(); err != nil {
		t.Error("Failed to create channel")
		t.Fatal(err)
	}

	channels, errChannels := models.GetChannels(senso.Address, service.Index)
	if errChannels != nil {
		t.Error("Failed to load back channels")
		t.Fatal(errChannels)
	}

	if channels[0].Index+1 != channels[1].Index {
		t.Fatal("index incr not working for channels")
	}

}
