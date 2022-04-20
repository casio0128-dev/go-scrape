package profile

import (
	"github.com/sclevine/agouti"
	"strings"
)

type TargetType string

const (
	IsCSSSelector = TargetType("SELECTOR")
	IsXPath       = TargetType("XPATH")
)

func (t TargetType) FindFunc(page *agouti.Page) func(string) *agouti.Selection {
	switch TargetType(strings.ToUpper(string(t))) {
	case IsXPath:
		return page.FindByXPath
	case IsCSSSelector:
		return page.Find
	}
	return page.Find
}
