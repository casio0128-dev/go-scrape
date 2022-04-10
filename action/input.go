package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type InputAction struct {
	name           string
	selector, text string
}

func NewInputAction(name string, selector string, text string) *InputAction {
	return &InputAction{name: name, selector: selector, text: text}
}

func (ia *InputAction) Name() string {
	return ia.name
}

func (ia *InputAction) Do(page *agouti.Page) error {
	if selection := page.Find(ia.selector); selection != nil {
		return selection.Fill(ia.text)
	}
	return NotExistsElement(ia.selector)
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
