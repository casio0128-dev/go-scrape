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
			return NewScreenShotAction(name, arg)
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
	case map[string]interface{}:
		var selector string
		if target, ok := arg[Target].(string); !ok {
			selector = target
		}
		if strings.EqualFold(selector, "") {
			selector = beforeSelector
		}

		switch name {
		case Input:
			if text, ok := arg[Text].(string); ok {
				a := NewInputAction(name, selector, text)
				return a
			}
		case SendKey:
			if keys, ok := arg[TypKey].(string); ok {
				return NewSendKeyAction(name, selector, keys)
			}
		case Select:
			if text, ok := arg[Text].(string); ok {
				return NewSelectAction(name, selector, text)
			}
		case Upload:
			if fileName, ok := arg[FileName].(string); ok {
				return NewSelectAction(name, selector, fileName)
			}
		case If:
			condMap := NewConditionMap()
			for conditionKey, actionValues := range arg {
				var acts []Action
				values, ok := actionValues.([]interface{})
				if !ok {
					return nil
				}

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
	fmt.Println("DEBUG: Not find action name.")
	return nil
}

func NotExistsElement(selector string) error {
	return fmt.Errorf("%s is not find element.\n", selector)
}

func NotActualFormat(name string) error {
	return fmt.Errorf("%s is invalid format\n", name)
}
