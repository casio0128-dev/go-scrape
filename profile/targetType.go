package profile

import "github.com/sclevine/agouti"

type TargetType string

const (
	IsCSSSelector = TargetType("Selector")
	IsXPath       = TargetType("Xpath")
)

func (t TargetType) FindFunc(page *agouti.Page) func(string) *agouti.Selection {
	switch t {
	case IsXPath:
		return page.FindByXPath
	case IsCSSSelector:
		return page.Find
	}
	return page.Find
}
