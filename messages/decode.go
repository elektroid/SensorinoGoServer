// This package contains code to convert json received from sensorino "base" to proper
// go structure.

package messages

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type BaseMessage struct {
	From      string   `json:"from"`
	To        string   `json:"to"`
	Type      string   `json:"type"`
	ServiceId []int64  `json:"serviceId"`
	DataType  []string `json:"dataType"`
	Count     []int64  `json:"count"`
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
	b.DataType = []string{}
	b.Count = []int64{}
	var rawData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonMsg), &rawData); err != nil {
		return nil, err
	}
	f, ok := rawData["from"]
	if !ok {
		return nil, errors.New("No from in base message")
	}

	switch reflect.ValueOf(f).Kind() {
	case reflect.String:
		b.From = f.(string)
	default:
		return nil, errors.New("Invalid from in base message")
	}
	delete(rawData, "from")

	to, ok := rawData["to"]
	if !ok {
		return nil, errors.New("No to in base message")
	}
	if reflect.ValueOf(to).Kind() != reflect.String {
		return nil, errors.New("Invalid to in base message")
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
	value := reflect.ValueOf(sIds)
	fmt.Printf("%+v\n", sIds)
	fmt.Println(value.Kind())
	if ok {
		switch value.Kind() {
		case reflect.Float32, reflect.Float64:
			b.ServiceId = []int64{int64(value.Float())}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			b.ServiceId = []int64{int64(value.Int())}
		case reflect.Slice:
			//we want either ints, floats or strings
			b.ServiceId = make([]int64, value.Len())
			for i := 0; i < value.Len(); i++ {

				//fmt.Println(value.Index(0).Int())
				switch value.Index(i).Elem().Kind() {
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					b.ServiceId[i] = int64(value.Index(i).Elem().Int())
				case reflect.String:
					v, err := strconv.ParseInt(value.Index(i).Elem().String(), 10, 64)
					if err != nil {
						return nil, errors.New("Failed to decode service id (atoi)")
					}
					b.ServiceId[i] = v
				}
				//	fmt.Printf("+%v\n", value.Index(i).Elem().Kind())
			}

		default:
			return nil, errors.New("Failed to decode service type, unexpected type")
		}
	}

	// could check type to see if it's normal not to find it
	delete(rawData, "serviceId")

	b.Data = rawData

	return b, nil

}

/*
func junk() {
	//we want either ints, floats or strings
	intValues := make([]int64, value.Len())
	floatValues := make([]float64, value.Len())
	stringValues := make([]string, value.Len())
	for i := 0; i < value.Len(); i++ {

		//fmt.Println(value.Index(0).Int())
		switch value.Index(i).Elem().Kind() {
		case reflect.Float32, reflect.Float64:
			floatValues = append(floatValues, float64(value.Index(i).Elem().Float()))
			if value.Len()-1 == i {
				b.ServiceId = floatValues
			}

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intValues = append(intValues, int64(value.Index(i).Elem().Int()))
			if value.Len()-1 == i {
				b.ServiceId = intValues
			}
		case reflect.String:
			stringValues = append(stringValues, value.Index(i).Elem().String())
			if value.Len()-1 == i {
				b.ServiceId = stringValues
			}
		}
		//	fmt.Printf("+%v\n", value.Index(i).Elem().Kind())
	}
}
*/
