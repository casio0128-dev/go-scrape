package action

import (
	"fmt"
	"github.com/sclevine/agouti"
	"go-scrape/profile"
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

	AssignText  = "assign-text"
	AssignTitle = "assign-title"
	AssignAttr  = "assign-attr"
	If          = "if"
)

const (
	Target   = "target"
	Text     = "text"
	TypKey   = "keys"
	FileName = "fileName"
	VarName  = "var"
	AttrName = "attr"
)

type Action interface {
	Name() string
	Do(*agouti.Page) error
	IsActual() bool
}

var beforeSelector string

func ParseAction(name string, prof *profile.Profile, args interface{}) Action {
	if strings.EqualFold(name, "") || len(name) <= 0 {
		return nil
	}

	switch arg := args.(type) {
	case string:
		switch name {
		case Click:
			if arg, err := parseVariables(arg, prof); err != nil {
				return nil
			} else {
				beforeSelector = arg
				return NewClickAction(name, arg)
			}
		case DoubleClick:
			if arg, err := parseVariables(arg, prof); err != nil {
				return nil
			} else {
				beforeSelector = arg
				return NewDoubleClickAction(name, arg)
			}
		case Wait:
			return NewWaitAction(name, arg)
		case ScreenShot:
			if arg, err := parseVariables(arg, prof); err != nil {
				return nil
			} else {
				return NewScreenShotAction(name, arg)
			}
		case To:
			if arg, err := parseVariables(arg, prof); err != nil {
				return nil
			} else {
				return NewToAction(name, arg)
			}
		case Cmd:
			if arg, err := parseVariables(arg, prof); err != nil {
				return nil
			} else {
				return NewCmdAction(name, arg)
			}
		case Reload:
			return NewReloadAction(name)
		case Exit:
			return NewExitAction(name)
		case Clear:
			return NewClearAction(name)
		case AssignTitle:
			return NewAssignTitleAction(name, arg, &prof.Variable)
		}
	case map[string]interface{}:
		var selector string
		if target, ok := arg[Target].(string); ok {
			selector = target
		}
		if strings.EqualFold(selector, "") {
			selector = beforeSelector
		}
		if selector, err := parseVariables(selector, prof); err != nil {
			return nil
		} else {
			selector = selector
		}

		switch name {
		case Input:
			if text, ok := arg[Text].(string); ok {
				if text, err := parseVariables(text, prof); err != nil {
					return nil
				} else {
					return NewInputAction(name, selector, text)
				}
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
				if fileName, err := parseVariables(fileName, prof); err != nil {
					return nil
				} else {
					return NewSelectAction(name, selector, fileName)
				}
			}
		case AssignText:
			if varName, ok := arg[VarName].(string); ok {
				return NewAssignTextAction(name, selector, varName, &prof.Variable)
			}
		case AssignAttr:
			if varName, ok := arg[VarName].(string); ok {
				if attrName, ok := arg[AttrName].(string); ok {
					return NewAssignAttrAction(name, selector, attrName, varName, &prof.Variable)
				}
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
							acts = append(acts, ParseAction(key, prof, val))
						}
					}
				}
				if parsedConditionKey, err := profile.Parse(conditionKey, prof.Variable); err == nil {
					condMap.Set(parsedConditionKey, acts)
				} else {
					continue
				}
			}
			return NewIfAction(name, condMap)
		}
	}
	fmt.Println("DEBUG: Not find action name.")
	return nil
}

func find(page *agouti.Page, prof *profile.Profile) func(string) *agouti.Selection {
	switch prof.TargetType {
	case profile.IsXPath:
		return page.FindByXPath
	case profile.IsCSSSelector:
		return page.Find
	}
	return page.Find
}

func parseVariables(str string, prof *profile.Profile) (string, error) {
	if parsed, err := profile.Parse(str, prof.Variable); err != nil {
		return "", err
	} else {
		return parsed, nil
	}
}

func NotExistsElement(selector string) error {
	return fmt.Errorf("%s is not find element.\n", selector)
}

func NotActualFormat(name string) error {
	return fmt.Errorf("%s is invalid format\n", name)
}
