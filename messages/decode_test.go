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
		&messages.BaseMessage{From: "0", To: "10", Type: "request", ServiceId: []int64{0}},
	},
	testDecode{
		`{ "from": "1.2.3.4", "to": "10", "type": "request", "serviceId": [ 0, "1", "2" ] }`,
		&messages.BaseMessage{From: "1.2.3.4", To: "10", Type: "request", ServiceId: []int64{0, 1, 2}},
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
			t.Logf("decoded does not produce proper result:\n %+v\n %+v", pair.Msg, b)
			j, _ := json.Marshal(b)
			t.Log(string(j))
			j2, _ := json.Marshal(b)
			t.Log(string(j2))

		}
	}
	if failures > 0 {
		t.Fatal("decoding failures:", failures)
	}

}
