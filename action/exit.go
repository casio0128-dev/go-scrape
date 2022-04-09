package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type ExitAction struct {
	name string
}

func NewExitAction(name string) *ExitAction {
	return &ExitAction{name: name}
}

func (ea *ExitAction) Name() string {
	return ea.name
}

func (ea *ExitAction) Do(page *agouti.Page) error {
	if ea.IsActual() {
		return page.Destroy()
	}
	return NotActualFormat(ea.name)
}

func (ea *ExitAction) IsActual() bool {
	if !strings.EqualFold(ea.name, "exit") {
		return false
	}
	return true
}
