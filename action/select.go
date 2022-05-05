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
	if sa.IsActual() {
		if selector, err := parseVariables(sa.selector, sa.prof); err != nil {
			return err
		} else {
			find := sa.prof.TargetType.FindFunc(page)
			if selection := find(selector); selection != nil {
				if text, err := parseVariables(sa.text, sa.prof); err != nil {
					return err
				} else {
					return selection.Select(text)
				}
			}
			return NotExistsElement(selector)
		}
	}
	return NotActualFormat(sa.Name())
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
