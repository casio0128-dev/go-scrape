package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type SelectAction struct {
	name           string
	selector, text string
}

func NewSelectAction(name string, selector string, text string) *SelectAction {
	return &SelectAction{name: name, selector: selector, text: text}
}

func (sa *SelectAction) Name() string {
	return sa.name
}

func (sa *SelectAction) Do(page *agouti.Page) error {
	if selection := page.Find(sa.selector); selection != nil {
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
