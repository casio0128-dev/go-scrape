package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type ClearAction struct {
	name string
	ok   bool
}

func NewClearAction(name string, ok bool) *ClearAction {
	return &ClearAction{name: name, ok: ok}
}

func (ca *ClearAction) Name() string {
	return ca.name
}

func (ca *ClearAction) Do(page *agouti.Page) error {
	if ca.IsActual() {
		if ca.ok {
			return page.ClearCookies()
		} else {
			return nil
		}
	}
	return NotActualFormat(ca.name)
}

func (ca *ClearAction) IsActual() bool {
	if !strings.EqualFold(ca.name, "clear") {
		return false
	}
	return true
}
