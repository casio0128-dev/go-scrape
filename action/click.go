package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type ClickAction struct {
	name     string
	selector string

	prof *profile.Profile
}

func NewClickAction(name, selector string, prof *profile.Profile) *ClickAction {
	return &ClickAction{name: name, selector: selector, prof: prof}
}

func (ca *ClickAction) Name() string {
	return ca.name
}

func (ca *ClickAction) Do(page *agouti.Page) error {
	if ca.IsActual() {
		find := ca.prof.TargetType.FindFunc(page)
		if selector, err := parseVariables(ca.selector, ca.prof); err != nil {
			return err
		} else {
			if selection := find(selector); selection != nil {
				return selection.Click()
			}
			return NotExistsElement(selector)
		}
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
