package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type ClearAction struct {
	name string
}

func NewClearAction(name string) *ClearAction {
	return &ClearAction{name: name}
}

func (ca *ClearAction) Name() string {
	return ca.name
}

func (ca *ClearAction) Do(page *agouti.Page) error {
	if ca.IsActual() {
		return page.ClearCookies()
	}
	return NotActualFormat(ca.name)
}

func (ca *ClearAction) IsActual() bool {
	if !strings.EqualFold(ca.name, "clear") {
		return false
	}
	return true
}
