package action

import (
	"fmt"
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type AssignTextAction struct {
	name     string
	selector *agouti.Selection

	key  string
	vars *profile.Variable
}

func NewAssignTextAction(name string, selector *agouti.Selection, key string, vars *profile.Variable) *AssignTextAction {
	return &AssignTextAction{name: name, selector: selector, key: key, vars: vars}
}

func (aa *AssignTextAction) Name() string {
	return aa.name
}

func (aa *AssignTextAction) Do(_ *agouti.Page) error {
	if aa.IsActual() {
		t, err := aa.selector.Text()
		if err != nil {
			return err
		}
		aa.vars.Set(aa.key, t)
	}
	return nil
}

func (aa *AssignTextAction) IsActual() bool {
	fmt.Println(aa.name, aa.key, aa.selector, *aa.vars)
	if !strings.EqualFold(aa.name, "assign-text") {
		return false
	}
	if aa.selector == nil {
		return false
	}
	if strings.EqualFold(aa.key, "") {
		return false
	}
	return true
}
