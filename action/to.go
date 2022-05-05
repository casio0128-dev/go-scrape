package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type ToAction struct {
	name string
	url  string

	prof *profile.Profile
}

func NewToAction(name string, url string, prof *profile.Profile) *ToAction {
	return &ToAction{name: name, url: url, prof: prof}
}

func (ta *ToAction) Name() string {
	return ta.name
}

func (ta *ToAction) Do(page *agouti.Page) error {
	if ta.IsActual() {
		if url, err := parseVariables(ta.url, ta.prof); err != nil {
			return err
		} else {
			return page.Navigate(url)
		}
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
