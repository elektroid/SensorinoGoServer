package events_test

import (
	"beezo/events"
	"testing"
	"time"
)

type dispatcher struct {
	t        *testing.T
	dispatch func(e events.Event)
	count    int
}

func NewTestDispatcher(t *testing.T, d func(e events.Event)) *dispatcher {
	return &dispatcher{t, d, 0}
}

func (this *dispatcher) Dispatch(e events.Event) {
	this.count++
	this.t.Log("upate count: ", this.count)
	this.dispatch(e)
}

func TestEvent(t *testing.T) {

	ev := events.Event{
		Type:             events.MissedSensorinoEvent,
		SensorinoAddress: "1.2.3.4",
		ServiceIndex:     0,
		ChannelIndex:     0,
	}

	disp := NewTestDispatcher(t, func(e events.Event) {
		if e.SensorinoAddress != ev.SensorinoAddress {
			t.Fatal("dispatched message has wrong address")
		}
		if e.ServiceIndex != ev.ServiceIndex {
			t.Fatal("wrong service index")
		}
		if e.Type != ev.Type {
			t.Fatal("wrong event type publish")
		}
	})

	events.RegisterDispatcher(disp)

	events.Publish(ev)
	events.Publish(ev)

	time.Sleep(10 * time.Millisecond)

	if disp.count != 2 {
		t.Fatal("missed event, count is ", disp.count)
	}
}
