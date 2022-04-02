package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type AssignAction struct {
	name     string
	selector string
	key      string
	vars     *profile.Variable
}

func (aa *AssignAction) Name() string {
	return aa.name
}

func (aa *AssignAction) Do(page *agouti.Page) error {
	if aa.IsActual() {
		if selector := page.Find(aa.selector); selector != nil {
			text, err := selector.Text()
			if err != nil {
				aa.vars.Set(aa.key, text)
			}
		}
	}
}

func (aa *AssignAction) IsActual() bool {
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
