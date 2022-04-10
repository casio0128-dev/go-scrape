package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type SendKeyAction struct {
	name          string
	selector, key string
}

func NewSendKeyAction(name string, selector *agouti.Selection, key string) *SendKeyAction {
	return &SendKeyAction{name: name, selector: selector, key: key}
}

func (ska *SendKeyAction) Name() string {
	return ska.name
}

func (ska *SendKeyAction) Do(page *agouti.Page) error {
	if selection := page.Find(ska.selector); selection != nil {
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
