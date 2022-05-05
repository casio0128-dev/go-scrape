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
	if dca.IsActual() {
		if selector, err := parseVariables(dca.selector, dca.prof); err != nil {
			return err
		} else {
			if selection := page.Find(selector); selection != nil {
				return selection.Click()
			}
			return NotExistsElement(selector)
		}
	}
	return NotActualFormat(dca.Name())
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
