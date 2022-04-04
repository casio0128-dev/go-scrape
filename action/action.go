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
	Target   = "target"
	Text     = "text"
	TypKey   = "keys"
	FileName = "fileName"
)

type Action interface {
	Name() string
	Do(*agouti.Page) error
	IsActual() bool
}

var beforeSelector string

func ParseAction(name string, args interface{}) Action {
	if strings.EqualFold(name, "") || len(name) <= 0 {
		return nil
	}

	switch arg := args.(type) {
	case string:
		switch name {
		case Click:
			beforeSelector = arg
			return NewClickAction(name, arg)
		case DoubleClick:
			beforeSelector = arg
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
		if strings.EqualFold(selector, "") {
			selector = beforeSelector
		}

		switch name {
		case Input:
			text := arg[Text]
			return NewInputAction(name, selector, text)
		case SendKey:
			keys := arg[TypKey]
			return NewSendKeyAction(name, selector, keys)
		case Select:
			text := arg[Text]
			return NewSelectAction(name, selector, text)
		case Upload:
			fileName := arg[FileName]
			return NewSelectAction(name, selector, fileName)
		}
	case map[string][]interface{}:
		switch name {
		case If:
			var condMap ConditionMap
			for conditionKey, values := range arg {
				var acts []Action
				for _, value := range values {
					if act, ok := value.(map[string]interface{}); ok {
						for key, val := range act {
							acts = append(acts, ParseAction(key, val))
						}
					}
				}
				condMap.Set(conditionKey, acts)
			}
			return NewIfAction(name, condMap)

		}
	}
	return nil
}

func NotExistsElement(selector string) error {
	return fmt.Errorf("%s is not find element.\n", selector)
}

func NotActualFormat(name string) error {
	return fmt.Errorf("%s is invalid format\n", name)
}
