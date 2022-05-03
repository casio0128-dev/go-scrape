package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type ReloadAction struct {
	name string
	ok   bool
}

func NewReloadAction(name string, ok bool) *ReloadAction {
	return &ReloadAction{name: name, ok: ok}
}

func (ra *ReloadAction) Name() string {
	return ra.name
}

func (ra *ReloadAction) Do(page *agouti.Page) error {
	if ra.IsActual() {
		if ra.ok {
			return page.Refresh()
		} else {
			return nil
		}
	}
	return NotActualFormat(ra.name)
}

func (ra *ReloadAction) IsActual() bool {
	if !strings.EqualFold(ra.name, "reload") {
		return false
	}
	return true
}
