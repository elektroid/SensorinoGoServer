// This package contains code to convert json received from sensorino "base" to proper
// go structure.

package messages

import (
	"encoding/json"
	"errors"
)

type BaseMessage struct {
	From      string   `json:"from"`
	To        string   `json:"to"`
	Type      string   `json:"type"`
	ServiceId []int    `json:"serviceId"`
	DataType  []string `json:"dataType"`
	Count     []int    `json:"count"`
	Data      map[string]interface{}
}

/*
	We would have liked to  directly unmarshal to our message structure but:
	- base serialization is not the safest
	- some fields are (or were) polymorphic (single elt vs array of )
	- published data comes in various shapes and random occurences
	All these aspects could be coded in a better way within the base
	firmware, but for now we have to cope with it
*/

func DecodeMessage(jsonMsg string) (*BaseMessage, error) {
	b := &BaseMessage{}
	var rawData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonMsg), &rawData); err != nil {
		return nil, err
	}
	f, ok := rawData["from"]
	if !ok {
		return nil, errors.New("No from in base message")
	}
	b.From = f.(string)
	delete(rawData, "from")

	to, ok := rawData["to"]
	if !ok {
		return nil, errors.New("No to in base message")
	}
	b.To = to.(string)
	delete(rawData, "to")

	mType, ok := rawData["type"]
	if !ok {
		return nil, errors.New("No type in base message")
	}
	b.Type = mType.(string)
	delete(rawData, "type")

	sIds, ok := rawData["serviceId"]
	if ok {
		b.ServiceId = sIds.([]int)
	}
	// could check type to see if it's normal not to find it
	delete(rawData, "serviceId")

	b.Data = rawData

	return b, nil

}
