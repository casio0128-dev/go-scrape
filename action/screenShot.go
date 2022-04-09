package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type ScreenShotAction struct {
	name     string
	fileName string
}

func NewScreenShotAction(name string, fileName string) *ScreenShotAction {
	return &ScreenShotAction{name: name, fileName: fileName}
}

func (ssa *ScreenShotAction) Name() string {
	return ssa.name
}

func (ssa *ScreenShotAction) Do(page *agouti.Page) error {
	if ssa.IsActual() {
		return page.Screenshot(ssa.fileName)
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
