package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type InputAction struct {
	name           string
	selector, text string
	prof           *profile.Profile
}

func NewInputAction(name string, selector string, text string, prof *profile.Profile) *InputAction {
	return &InputAction{name: name, selector: selector, text: text, prof: prof}
}

func (ia *InputAction) Name() string {
	return ia.name
}

func (ia *InputAction) Do(page *agouti.Page) error {
	if ia.IsActual() {
		if selector, err := parseVariables(ia.selector, ia.prof); err != nil {
			return err
		} else {
			find := ia.prof.TargetType.FindFunc(page)
			if selection := find(selector); selection != nil {
				return selection.Fill(ia.text)
			}
			return NotExistsElement(selector)
		}
	}
	return NotActualFormat(ia.Name())
}

func (ia *InputAction) IsActual() bool {
	if !strings.EqualFold(ia.name, "input") {
		return false
	}
	if strings.EqualFold(ia.selector, "") {
		return false
	}
	return true
}
