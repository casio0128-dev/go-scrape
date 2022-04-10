package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type DoubleClickAction struct {
	name     string
	selector *agouti.Selection
}

func NewDoubleClickAction(name string, selector *agouti.Selection) *DoubleClickAction {
	return &DoubleClickAction{name: name, selector: selector}
}

func (dca *DoubleClickAction) Name() string {
	return dca.name
}

func (dca *DoubleClickAction) Do(_ *agouti.Page) error {
	if dca.IsActual() {
		return dca.selector.DoubleClick()
	}
	return NotExistsElement(dca.selector.String())
}

func (dca *DoubleClickAction) IsActual() bool {
	if !strings.EqualFold(dca.name, "doubleClick") {
		return false
	}
	if dca.selector == nil {
		return false
	}
	return true
}
