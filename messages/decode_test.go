package messages_test

import (
	"beezo/messages"
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
		&messages.BaseMessage{From: "0", To: "10", Type: "request", ServiceId: []int{0}},
	},
	testDecode{
		"{ \"from\": 0, \"to\": 10, \"type\": \"request\", \"serviceId\": [ \"0\", \"1\", \"2\" ] }",
		&messages.BaseMessage{From: "0", To: "10", Type: "request", ServiceId: []int{0, 1, 2}},
	},
}

func TestDecode(t *testing.T) {

	for _, pair := range set {
		b, err := messages.DecodeMessage(pair.Json)
		if err != nil {
			t.Fatal("failed to decode", err.Error())
		}
		if !reflect.DeepEqual(b, pair.Msg) {
			t.Fatal("decoded does not produce proper result", pair.Msg, b)
		}

	}

}
