package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type AssignAttrAction struct {
	name     string
	selector string
	attrName string

	key  string
	vars *profile.Variable
}

func NewAssignAttrAction(name string, selector string, attrName string, key string, vars *profile.Variable) *AssignAttrAction {
	return &AssignAttrAction{name: name, selector: selector, attrName: attrName, key: key, vars: vars}
}

func (aa *AssignAttrAction) Name() string {
	return aa.name
}

func (aa *AssignAttrAction) Do(page *agouti.Page) error {
	if aa.IsActual() {
		if selector := page.Find(aa.selector); selector != nil {
			a, err := selector.Attribute(aa.attrName)
			if err != nil {
				return err
			}
			aa.vars.Set(aa.key, a)
		}
	}
	return nil
}

func (aa *AssignAttrAction) IsActual() bool {
	if !strings.EqualFold(aa.name, "assign-attr") {
		return false
	}
	if strings.EqualFold(aa.selector, "") {
		return false
	}
	if strings.EqualFold(aa.key, "") {
		return false
	}
	if strings.EqualFold(aa.attrName, "") {
		return false
	}
	return true
}
