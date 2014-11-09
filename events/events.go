package events

// we want to publish events for different systems
// mqtt websocket can be used to make a rich UI with live events / updates
// as central/base node for sensorinowe also want to take some actions
//
// The point of using golang chan here would be to treat events in different / concurrent place and not block
// is it worth it ? is it clever ?

const (
	// within server stuff happens
	NewSensorinoEvent    = iota
	NewServiceEvent      = iota
	MissedSensorinoEvent = iota
	MissedServiceEvent   = iota
	MissedChanEvent      = iota
)

type Event struct {
	Type             int
	SensorinoAddress string
	ServiceIndex     int64
	ChannelIndex     int64
}

type Dispatcher interface {
	Dispatch(Event)
}

var dispatchers = make([]Dispatcher, 0, 2)

func Publish(event Event) {
	for _, dispatcher := range dispatchers {
		go func() { dispatcher.Dispatch(event) }()
	}
}

func RegisterDispatcher(d Dispatcher) {
	dispatchers = append(dispatchers, d)
}

func MissedSensorino(address string) {
	Publish(Event{Type: MissedSensorinoEvent, SensorinoAddress: address})
}
