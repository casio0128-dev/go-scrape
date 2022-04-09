package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type ReloadAction struct {
	name string
}

func NewReloadAction(name string) *ReloadAction {
	return &ReloadAction{name: name}
}

func (ra *ReloadAction) Name() string {
	return ra.name
}

func (ra *ReloadAction) Do(page *agouti.Page) error {
	return page.Refresh()
}

func (ra *ReloadAction) IsActual() bool {
	if !strings.EqualFold(ra.name, "reload") {
		return false
	}
	return true
}
