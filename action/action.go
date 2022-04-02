package action

import (
	"fmt"
	"github.com/sclevine/agouti"
	"strings"
)

const (
	Click       = "click"
	DoubleClick = "doubleClick"
	Wait        = "wait"
	ScreenShot  = "screenShot"
	To          = "to"
	Reload      = "reload"
	Exit        = "exit"
	Cmd         = "cmd"
	Clear       = "clear"

	Input   = "input"
	SendKey = "sendKey"
	Select  = "select"
	Upload  = "upload"

	Assign = "assign"
	If     = "if"
)

const (
	Target = "target"
	Text   = "text"
)

type Action interface {
	Name() string
	Do(*agouti.Page) error
	IsActual() bool
}

func sample() {
	s := ParseAction(Click, "body")
}

func ParseAction(name string, args interface{}) Action {
	if strings.EqualFold(name, "") || len(name) <= 0 {
		return nil
	}

	switch arg := args.(type) {
	case string:
		switch name {
		case Click:
			return NewClickAction(name, arg)
		case DoubleClick:
			return NewDoubleClickAction(name, arg)
		case Wait:
			return NewWaitAction(name, arg)
		case ScreenShot:
			return NewWaitAction(name, arg)
		case To:
			return NewToAction(name, arg)
		case Cmd:
			return NewCmdAction(name, arg)
		case Reload:
			return NewReloadAction(name)
		case Exit:
			return NewExitAction(name)
		case Clear:
			return NewClearAction(name)
		}
	case map[string]string:
		selector := arg[Target]

		switch name {
		case Input:
			return NewInputAction(name, "selector not specify by before selector in click or double click element.", txt)
		case SendKey:
		case Select:
		case Upload:
		}
	}
	return nil
}

func NotExistsElement(selector string) error {
	return fmt.Errorf("%s is not find element.\n")
}

func NotActualFormat(name string) error {
	return fmt.Errorf("%s is invalid format\n", name)
}
