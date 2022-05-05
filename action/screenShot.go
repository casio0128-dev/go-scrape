package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type ScreenShotAction struct {
	name     string
	fileName string

	prof *profile.Profile
}

func NewScreenShotAction(name string, fileName string, prof *profile.Profile) *ScreenShotAction {
	return &ScreenShotAction{name: name, fileName: fileName, prof: prof}
}

func (ssa *ScreenShotAction) Name() string {
	return ssa.name
}

func (ssa *ScreenShotAction) Do(page *agouti.Page) error {
	if ssa.IsActual() {
		if name, err := parseVariables(ssa.fileName, ssa.prof); err != nil {
			return err
		} else {
			return page.Screenshot(name)
		}
	}
	return NotActualFormat(ssa.name)
}

func (ssa *ScreenShotAction) IsActual() bool {
	if !strings.EqualFold(ssa.name, "screenShot") {
		return false
	}
	if strings.EqualFold(ssa.fileName, "") {
		return false
	}
	return true
}
