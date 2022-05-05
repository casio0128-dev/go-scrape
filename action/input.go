package action

import (
	"fmt"
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
				if text, err := parseVariables(ia.text, ia.prof); err != nil {
					return err
				} else {
					fmt.Println(text, ia.prof)
					return selection.Fill(text)
				}
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
