package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type ClickAction struct {
	name     string
	selector string
}

func NewClickAction(name, selector string) *ClickAction {
	return &ClickAction{name: name, selector: selector}
}

func (ca *ClickAction) Name() string {
	return ca.name
}

func (ca *ClickAction) Do(page *agouti.Page) error {
	if selection := page.Find(ca.selector); selection != nil {
		return selection.Click()
	}
	return NotExistsElement(ca.selector)
}

func (ca *ClickAction) IsActual() bool {
	if !strings.EqualFold(ca.name, "click") {
		return false
	}
	if strings.EqualFold(ca.selector, "") {
		return false
	}
	return true
}
