package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type AssignTitleAction struct {
	name string
	key  string
	vars *profile.Variable
}

func NewAssignTitleAction(name string, key string, vars *profile.Variable) *AssignTitleAction {
	return &AssignTitleAction{name: name, key: key, vars: vars}
}

func (aa *AssignTitleAction) Name() string {
	return aa.name
}

func (aa *AssignTitleAction) Do(page *agouti.Page) error {
	if aa.IsActual() {
		t, err := page.Title()
		if err != nil {
			return err
		}
		aa.vars.Set(aa.key, t)
	}
	return nil
}

func (aa *AssignTitleAction) IsActual() bool {
	if !strings.EqualFold(aa.name, "assign-title") {
		return false
	}
	if strings.EqualFold(aa.key, "") {
		return false
	}
	return true
}
