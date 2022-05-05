package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type SendKeyAction struct {
	name          string
	selector, key string

	prof *profile.Profile
}

func NewSendKeyAction(name string, selector string, key string, prof *profile.Profile) *SendKeyAction {
	return &SendKeyAction{name: name, selector: selector, key: key, prof: prof}
}

func (ska *SendKeyAction) Name() string {
	return ska.name
}

func (ska *SendKeyAction) Do(page *agouti.Page) error {
	if ska.IsActual() {
		if selector, err := parseVariables(ska.selector, ska.prof); err != nil {
			return err
		} else {
			find := ska.prof.TargetType.FindFunc(page)
			if selection := find(selector); selection != nil {
				if pressKey, err := parseVariables(ska.key, ska.prof); err != nil {
					return err
				} else {
					return selection.SendKeys(pressKey)
				}
			}
		}
		return NotExistsElement(ska.selector)
	}
	return NotActualFormat(ska.Name())
}

func (ska *SendKeyAction) IsActual() bool {
	if !strings.EqualFold(ska.name, "sendKey") {
		return false
	}
	if strings.EqualFold(ska.selector, "") {
		return false
	}
	return true
}
