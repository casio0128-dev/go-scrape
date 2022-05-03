package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/profile"
	"strconv"
	"strings"
)

type ForAction struct {
	name       string
	start, end int

	operations []Action

	indexVar string
	prof     *profile.Profile
}

func (fa *ForAction) Name() string {
	return fa.name
}

func (fa *ForAction) Do(page *agouti.Page) error {
	if fa.IsActual() {
		prof := fa.prof

		for i := fa.start; i < fa.end; fa.next(&i) {
			prof.Variable.Set(fa.indexVar, strconv.Itoa(i))
			for _, operation := range fa.operations {
				if err := operation.Do(page); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (fa *ForAction) next(current *int) {
	if fa.start <= fa.end {
		*current += 1
	} else if fa.start > fa.end {
		*current -= 1
	}
}

func (fa *ForAction) IsActual() bool {
	if !strings.EqualFold(fa.name, "for") {
		return false
	}

	return true
}
