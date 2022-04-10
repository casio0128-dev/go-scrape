package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type AssignAction struct {
	name       string
	selector   *agouti.Selection
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
		var (
			content string
			err     error
		)

		switch aa.targetType {
		case TargetTypeText:
			if content, err = aa.selector.Text(); err != nil {
				return err
			}
		case TargetTypeAttribute:
			if content, err = aa.selector.Attribute(aa.attrName); err != nil {
				return err
			}
		case TargetTypeTitle:
			if content, err = page.Title(); err != nil {
				return err
			}
		}
		aa.vars.Set(aa.key, content)
	}
	return nil
}

func (aa *AssignAction) IsActual() bool {
	if !strings.EqualFold(aa.name, "assign") {
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
