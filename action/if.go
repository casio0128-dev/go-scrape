package action

import (
	"github.com/sclevine/agouti"
	"go-scrape/common"
	"regexp"
	"strconv"
	"strings"
)

const (
	_EQUAL        = "=="
	_NOT_EQUAL    = "!="
	_GREATER_THAN = ">"
	_LESS_THAN    = "<"
)

// TODO: Condition関連のみファイルを分割する？
type ConditionMap map[Condition][]Action

func (c ConditionMap) GetConditions() []Condition {
	var conditions []Condition
	for condition, _ := range c {
		conditions = append(conditions, condition)
	}
	return conditions
}

func (c ConditionMap) Get(key string) []Action {
	return c[Condition(key)]
}

func (c ConditionMap) Set(key string, value []Action) {
	c[Condition(key)] = value
}

type Condition string

const REGEX_PATTERN = `[==|!=|>|<|]`

func (c Condition) Expr() bool {
	regex, err := regexp.Compile(REGEX_PATTERN)
	if err != nil {
		return false
	}
	operator := regex.FindStringSubmatch(string(c))
	formula := regex.Split(string(c), -1)
	switch operator[0] {
	case _EQUAL:
		return c.equal(formula[0], formula[1])
	case _NOT_EQUAL:
		return c.notEqual(formula[0], formula[1])
	case _GREATER_THAN:
		return c.greaterThan(formula[0], formula[1])
	case _LESS_THAN:
		return c.lessThan(formula[0], formula[1])
	}
	return false
}

func (c Condition) equal(left, right string) bool {
	return strings.EqualFold(left, right)
}

func (c Condition) notEqual(left, right string) bool {
	return !c.equal(left, right)
}

func (c Condition) greaterThan(left, right string) bool {
	if common.IsInt(left) && common.IsInt(right) {
		l, _ := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)
		return l > r
	} else {
		return len(left) > len(right)
	}
}

func (c Condition) lessThan(left, right string) bool {
	if common.IsInt(left) && common.IsInt(right) {
		l, _ := strconv.Atoi(left)
		r, _ := strconv.Atoi(right)
		return l < r
	} else {
		return len(left) < len(right)
	}
}

type IfAction struct {
	name string
	proc ConditionMap
}

func NewIfAction(name string, proc ConditionMap) *IfAction {
	return &IfAction{name: name, proc: proc}
}

func (ia *IfAction) Name() string {
	return ia.name
}

func (ia *IfAction) Do(page *agouti.Page) error {
	if ia.IsActual() {
		for _, condition := range ia.proc.GetConditions() {
			if condition.Expr() {
				for _, act := range ia.proc[condition] {
					if err := act.Do(page); err != nil {
						return err
					}
				}
				break
			}
		}
	}
	return nil
}

func (ia *IfAction) IsActual() bool {
	if !strings.EqualFold(ia.name, "if") {
		return false
	}
	if len(ia.proc) <= 0 {
		return false
	}
	return true
}
