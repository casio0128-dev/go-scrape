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
	find := ska.prof.TargetType.FindFunc(page)
	if selection := find(ska.selector); selection != nil {
		return selection.SendKeys(ska.key)
	}
	return NotExistsElement(ska.selector)
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
