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
		if err := page.Session().Refresh(); err != nil {
			return err
		}
		if err := page.ClearCookies(); err != nil {
			return err
		}
		return nil
	}
	return NotActualFormat(ca.name)
}

func (ca *ClearAction) IsActual() bool {
	if !strings.EqualFold(ca.name, "clear") {
		return false
	}
	return true
}
