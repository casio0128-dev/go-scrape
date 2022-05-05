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
	prof *profile.Profile
}

func NewAssignTextAction(name string, selector string, key string, prof *profile.Profile) *AssignTextAction {
	return &AssignTextAction{name: name, selector: selector, key: key, prof: prof}
}

func (aa *AssignTextAction) Name() string {
	return aa.name
}

func (aa *AssignTextAction) Do(page *agouti.Page) error {
	if aa.IsActual() {
		vars := &(aa.prof.Variable)
		find := aa.prof.TargetType.FindFunc(page)
		if selector, err := parseVariables(aa.selector, aa.prof); err != nil {
			return err
		} else {
			if selection := find(selector); selection != nil {
				t, err := selection.Text()
				if err != nil {
					return err
				}
				vars.Set(aa.key, t)
				return nil
			}
			return NotExistsElement(selector)
		}
	}
	return nil
}

func (aa *AssignTextAction) IsActual() bool {
	if !strings.EqualFold(aa.name, "assign-text") {
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
