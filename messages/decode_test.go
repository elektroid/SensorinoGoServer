package messages_test

import (
	"beezo/messages"
	"encoding/json"
	"reflect"
	"testing"
)

type testDecode struct {
	Json string
	Msg  *messages.BaseMessage
}

var set = []testDecode{
	testDecode{
		`{ "from": "0", "to": "10", "type": "request", "serviceId": 0 }`,
		&messages.BaseMessage{From: "0", To: "10", Type: messages.Request, ServiceId: []int64{0}},
	},
	testDecode{
		`{ "from": "1.2.3.4", "to": "10", "type": "request", "serviceId": [ 0, "1", "2" ] }`,
		&messages.BaseMessage{From: "1.2.3.4", To: "10", Type: messages.Request, ServiceId: []int64{0, 1, 2}},
	},

	testDecode{
		`{ "from": "1.2.3.4", "to": "10", "type": "publish", "serviceId":  2,  "switch": false  }`,
		&messages.BaseMessage{From: "1.2.3.4", To: "10", Type: messages.Publish, ServiceId: []int64{2}, Data: map[string]interface{}{"switch": false}},
	},

	testDecode{
		`{ "from": "10.2", "to": "0", "type": "publish", "serviceId": 2, "dataType": "Switch", "count": [ 0, 1 ] }`,
		&messages.BaseMessage{From: "10.2", To: "0", Type: messages.Publish, DataType: []string{"Switch"}, ServiceId: []int64{2}, Count: []int64{0, 1}},
	},
}

func TestDecode(t *testing.T) {

	failures := 0
	for _, pair := range set {
		b, err := messages.DecodeMessage(pair.Json)
		if err != nil {
			t.Fatal("failed to decode", err.Error())
		}
		if reflect.DeepEqual(b, pair.Msg) == false {
			failures++
			t.Logf("decoded does not produce proper result:\njson:%s\n %+v\n %+v", pair.Json, pair.Msg, b)
			j2, _ := json.Marshal(pair.Msg)
			t.Log(string(j2))

			j, _ := json.Marshal(b)
			t.Log(string(j))

		}
	}
	if failures > 0 {
		t.Fatal("decoding failures:", failures)
	}

}
