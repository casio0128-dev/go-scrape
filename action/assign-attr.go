package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type AssignAttrAction struct {
	name     string
	selector *agouti.Selection
	attrName string

	key  string
	vars *profile.Variable
}

func NewAssignAttrAction(name string, selector *agouti.Selection, attrName string, key string, vars *profile.Variable) *AssignAttrAction {
	return &AssignAttrAction{name: name, selector: selector, attrName: attrName, key: key, vars: vars}
}

func (aa *AssignAttrAction) Name() string {
	return aa.name
}

func (aa *AssignAttrAction) Do(_ *agouti.Page) error {
	if aa.IsActual() {
		a, err := aa.selector.Attribute(aa.attrName)
		if err != nil {
			return err
		}
		aa.vars.Set(aa.key, a)
	}
	return nil
}

func (aa *AssignAttrAction) IsActual() bool {
	if !strings.EqualFold(aa.name, "assign-attr") {
		return false
	}
	if aa.selector == nil {
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
