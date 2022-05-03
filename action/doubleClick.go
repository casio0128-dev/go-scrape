package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type DoubleClickAction struct {
	name     string
	selector string
	prof     *profile.Profile
}

func NewDoubleClickAction(name string, selector string, prof *profile.Profile) *DoubleClickAction {
	return &DoubleClickAction{name: name, selector: selector, prof: prof}
}

func (dca *DoubleClickAction) Name() string {
	return dca.name
}

func (dca *DoubleClickAction) Do(page *agouti.Page) error {
	if selection := page.Find(dca.selector); selection != nil {
		return selection.Click()
	}
	return NotExistsElement(dca.selector)
}

func (dca *DoubleClickAction) IsActual() bool {
	if !strings.EqualFold(dca.name, "doubleClick") {
		return false
	}
	if strings.EqualFold(dca.selector, "") {
		return false
	}
	return true
}
