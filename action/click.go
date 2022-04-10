package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type ClickAction struct {
	name     string
	selector *agouti.Selection
}

func NewClickAction(name string, selector *agouti.Selection) *ClickAction {
	return &ClickAction{name: name, selector: selector}
}

func (ca *ClickAction) Name() string {
	return ca.name
}

func (ca *ClickAction) Do(_ *agouti.Page) error {
	if ca.IsActual() {
		return ca.selector.Click()
	}
	return NotExistsElement(ca.selector.Selectors().String())
}

func (ca *ClickAction) IsActual() bool {
	if !strings.EqualFold(ca.name, "click") {
		return false
	}
	if ca.selector == nil {
		return false
	}
	return true
}
