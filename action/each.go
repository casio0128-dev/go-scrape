package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strconv"
	"strings"
)

type EachAction struct {
	name       string
	target     string
	operations []Action

	indexVar string
	prof     *profile.Profile
}

func NewEachAction(name string, target string, operations []Action, indexVar string, prof *profile.Profile) *EachAction {
	return &EachAction{name: name, target: target, operations: operations, indexVar: indexVar, prof: prof}
}

func (ea *EachAction) Name() string {
	return ea.name
}

func (ea *EachAction) Do(page *agouti.Page) error {
	if ea.IsActual() {
		prof := ea.prof
		find := prof.TargetType.FindFunc(page)
		elements, err := find(ea.target).Count()
		if err != nil {
			return err
		}

		for i := 1; i <= elements; i++ {
			prof.Variable.Set(ea.indexVar, strconv.Itoa(i))
			for _, operation := range ea.operations {
				if err := operation.Do(page); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (ea *EachAction) IsActual() bool {
	if !strings.EqualFold(ea.name, "each") {
		return false
	}
	if strings.EqualFold(ea.target, "") {
		return false
	}
	return true
}
