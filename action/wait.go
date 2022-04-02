package action

import (
	"github.com/sclevine/agouti"
	"strings"
	"time"
)

type WaitAction struct {
	name     string
	waitTime string
}

func NewWaitAction(name string, waitTime string) *WaitAction {
	return &WaitAction{name: name, waitTime: waitTime}
}

func (wa *WaitAction) Name() string {
	return wa.name
}

func (wa *WaitAction) Do(_ *agouti.Page) error {
	if d, err := time.ParseDuration(wa.waitTime); err != nil {
		return err
	} else {
		<-time.After(d)
	}
	return nil
}

func (wa *WaitAction) IsActual() bool {
	if !strings.EqualFold(wa.name, "wait") {
		return false
	}
	if strings.EqualFold(wa.waitTime, "") {
		return false
	}
	return true
}
