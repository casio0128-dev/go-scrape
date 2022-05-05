package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strings"
)

type UploadAction struct {
	name     string
	selector string
	fileName string

	prof *profile.Profile
}

func NewUploadAction(name string, selector string, fileName string, prof *profile.Profile) *UploadAction {
	return &UploadAction{name: name, selector: selector, fileName: fileName, prof: prof}
}

func (ua *UploadAction) Name() string {
	return ua.name
}

func (ua *UploadAction) Do(page *agouti.Page) error {
	if ua.IsActual() {
		if selector, err := parseVariables(ua.selector, ua.prof); err != nil {
			return err
		} else {
			find := ua.prof.TargetType.FindFunc(page)
			if selection := find(selector); selection != nil {
				if fileName, err := parseVariables(ua.fileName, ua.prof); err != nil {
					return err
				} else {
					return selection.UploadFile(fileName)
				}
			}
			return NotExistsElement(selector)
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
