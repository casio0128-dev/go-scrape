package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type AssignTextAction struct {
	name     string
	selector string

	key  string
	vars *profile.Variable
}

func NewAssignTextAction(name string, selector string, key string, vars *profile.Variable) *AssignTextAction {
	return &AssignTextAction{name: name, selector: selector, key: key, vars: vars}
}

func (aa *AssignTextAction) Name() string {
	return aa.name
}

func (aa *AssignTextAction) Do(page *agouti.Page) error {
	if aa.IsActual() {
		if selector := page.Find(aa.selector); selector != nil {
			t, err := selector.Text()
			if err != nil {
				return err
			}
			aa.vars.Set(aa.key, t)
		}
	}
	return nil
}

func (aa *AssignTextAction) IsActual() bool {
	if !strings.EqualFold(aa.name, "assign") {
		return false
	}
	if strings.EqualFold(aa.selector, "") {
		return false
	}
	if strings.EqualFold(aa.key, "") {
		return false
	}
	return true
}
