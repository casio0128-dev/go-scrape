package action

import (
	"github.com/sclevine/agouti"
	"strings"
)

type UploadAction struct {
	name     string
	selector string
	fileName string
}

func NewUploadAction(name string, selector string, fileName string) *UploadAction {
	return &UploadAction{name: name, selector: selector, fileName: fileName}
}

func (ua *UploadAction) Name() string {
	return ua.name
}

func (ua *UploadAction) Do(page *agouti.Page) error {
	if ua.IsActual() {
		if selection := page.Find(ua.selector); selection != nil {
			return selection.UploadFile(ua.fileName)
		} else {
			return NotExistsElement(ua.selector)
		}
	}
	return NotActualFormat(ua.name)
}

func (ua *UploadAction) IsActual() bool {
	if !strings.EqualFold(ua.name, "upload") {
		return false
	}
	if strings.EqualFold(ua.selector, "") {
		return false
	}
	if strings.EqualFold(ua.fileName, "") {
		return false
	}
	return true
}
