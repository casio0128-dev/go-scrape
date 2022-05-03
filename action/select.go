package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type SelectAction struct {
	name           string
	selector, text string

	prof *profile.Profile
}

func NewSelectAction(name string, selector string, text string, prof *profile.Profile) *SelectAction {
	return &SelectAction{name: name, selector: selector, text: text, prof: prof}
}

func (sa *SelectAction) Name() string {
	return sa.name
}

func (sa *SelectAction) Do(page *agouti.Page) error {
	find := sa.prof.TargetType.FindFunc(page)
	if selection := find(sa.selector); selection != nil {
		return selection.Select(sa.text)
	}
	return NotExistsElement(sa.selector)
}

func (sa *SelectAction) IsActual() bool {
	if !strings.EqualFold(sa.name, "input") {
		return false
	}
	if strings.EqualFold(sa.selector, "") {
		return false
	}
	return true
}
