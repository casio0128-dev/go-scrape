package profile

import (
	"github.com/sclevine/agouti"
	"strings"
)

type BrowserOption struct {
	IsHeadless bool         `json:"isHeadless"`
	Data       string       `json:"userData"`
	Size       *BrowserSize `json:"size"`
}

type BrowserSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func (bo BrowserOption) ChromeOption() agouti.Option {
	var arguments = []string{}
	
	if bo.IsHeadless {
		arguments = append(arguments, isHeadless())
	}
	if bo.Size != nil {
		var width, height = 1280, 1280
		if bo.Size.Width > 0 {
			width = bo.Size.Width
		}
		if bo.Size.Height > 0 {
			height = bo.Size.Height
		}
		arguments = append(arguments, sizeBy(width, height))
	}
	if !strings.EqualFold(bo.Data, "") {
		arguments = append(arguments, userDataBy(bo.Data))
	}

	return agouti.ChromeOptions("args", arguments)
}
