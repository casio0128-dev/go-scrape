package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type ExitAction struct {
	name string
	ok   bool
}

func NewExitAction(name string, ok bool) *ExitAction {
	return &ExitAction{name: name, ok: ok}
}

func (ea *ExitAction) Name() string {
	return ea.name
}

func (ea *ExitAction) Do(page *agouti.Page) error {
	if ea.IsActual() {
		if ea.ok {
			return page.Destroy()
		} else {
			return nil
		}
	}
	return NotActualFormat(ea.name)
}

func (ea *ExitAction) IsActual() bool {
	if !strings.EqualFold(ea.name, "exit") {
		return false
	}
	return true
}
