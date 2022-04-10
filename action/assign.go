package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type AssignAction struct {
	name       string
	selector   string
	targetType string
	attrName   string

	key  string
	vars *profile.Variable
}

const (
	TargetTypeText      = "text"
	TargetTypeTitle     = "title"
	TargetTypeAttribute = "attribute"
)

func (aa *AssignAction) Name() string {
	return aa.name
}

func (aa *AssignAction) Do(page *agouti.Page) error {
	if aa.IsActual() {
		if selector := page.Find(aa.selector); selector != nil {
			var (
				content string
				err     error
			)

			switch aa.targetType {
			case TargetTypeText:
				if content, err = selector.Text(); err != nil {
					return err
				}
			case TargetTypeAttribute:
				if content, err = selector.Attribute(aa.attrName); err != nil {
					return err
				}
			case TargetTypeTitle:
				if content, err = page.Title(); err != nil {
					return err
				}
			}
			aa.vars.Set(aa.key, content)
		}
	}
	return nil
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
