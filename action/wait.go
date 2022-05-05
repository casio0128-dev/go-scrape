package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
	"time"
)

type WaitAction struct {
	name     string
	waitTime string
	prof     *profile.Profile
}

func NewWaitAction(name string, waitTime string, prof *profile.Profile) *WaitAction {
	return &WaitAction{name: name, waitTime: waitTime, prof: prof}
}

func (wa *WaitAction) Name() string {
	return wa.name
}

func (wa *WaitAction) Do(_ *agouti.Page) error {
	if wa.IsActual() {
		if waitTime, err := parseVariables(wa.waitTime, wa.prof); err != nil {
			return err
		} else {
			if d, err := time.ParseDuration(waitTime); err != nil {
				return err
			} else {
				<-time.After(d)
				return nil
			}
		}
	}
	return NotActualFormat(wa.Name())
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
