package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type ToAction struct {
	name string
	url  string
}

func NewToAction(name string, url string) *ToAction {
	return &ToAction{name: name, url: url}
}

func (ta *ToAction) Name() string {
	return ta.name
}

func (ta *ToAction) Do(page *agouti.Page) error {
	if ta.IsActual() {
		return page.Navigate(ta.url)
	}
	return NotActualFormat(ta.name)
}

func (ta *ToAction) IsActual() bool {
	if !strings.EqualFold(ta.name, "to") {
		return false
	}
	if strings.EqualFold(ta.url, "") ||
		(!strings.HasPrefix(ta.url, "http://") && !strings.HasPrefix(ta.url, "https://")) {
		return false
	}
	return true
}
