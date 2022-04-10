package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type UploadAction struct {
	name     string
	selector *agouti.Selection
	fileName string
}

func NewUploadAction(name string, selector *agouti.Selection, fileName string) *UploadAction {
	return &UploadAction{name: name, selector: selector, fileName: fileName}
}

func (ua *UploadAction) Name() string {
	return ua.name
}

func (ua *UploadAction) Do(page *agouti.Page) error {
	if ua.IsActual() {
		return ua.selector.UploadFile(ua.fileName)
	}
	return NotExistsElement(ua.selector.String())
}

func (ua *UploadAction) IsActual() bool {
	if !strings.EqualFold(ua.name, "upload") {
		return false
	}
	if ua.selector == nil {
		return false
	}
	if strings.EqualFold(ua.fileName, "") {
		return false
	}
	return true
}
