package events

import (
	"fmt"
	"log"
	"os"
)

type DebugDispatcher struct {
	logger *log.Logger
}

func NewDebugDispatcher() *DebugDispatcher {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	return &DebugDispatcher{logger}
}

func (this *DebugDispatcher) Start() {
	RegisterDispatcher(this)
	this.logger.Println("DebugDispatcher started")
}

func (this DebugDispatcher) Dispatch(e Event) {
	text := fmt.Sprintf("%+v", e)
	this.logger.Println(text)
}
