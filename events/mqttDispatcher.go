package events

import (
	"fmt"

	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

const (
	LocalMqttServer = "tcp://127.0.0.1:1883"
	ClientId        = "beezoServer"
	Channel         = "sensorino"
)

type MqttDispatcher struct {
	Broker   string
	ClientId string
	Channel  string
	c        *MQTT.MqttClient
}

var f MQTT.MessageHandler = func(client *MQTT.MqttClient, msg MQTT.Message) {
	fmt.Printf("Received message TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func (this MqttDispatcher) Start(broker string) {

	this.Broker = broker
	this.ClientId = ClientId
	this.Channel = Channel

	opts := MQTT.NewClientOptions().AddBroker(this.Broker).SetClientId(this.ClientId)
	opts.SetDefaultPublishHandler(f)

	this.c = MQTT.NewClient(opts)
	_, err := this.c.Start()
	if err != nil {
		panic(err)
	}

}

func (this MqttDispatcher) Dispatch(e Event) {
	text := fmt.Sprintf("%+v", e)
	receipt := this.c.Publish(MQTT.QOS_ONE, this.Channel, []byte(text))
	<-receipt
}
