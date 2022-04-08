package profile

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Profiles []Profile

type Profile struct {
	Name      string    `json:"name"`
	Variable  Variable  `json:"var"`
	Operation Operation `json:"operation"`
}

func Parse(str string, vars Variable) (string, error) {
	result := str
	reg, err := regexp.Compile(`{([a-z|A-Z]+)}`)
	if err != nil {
		return "", err
	}
	matches := reg.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		key := match[1]

		if !vars.IsExists(key) {
			continue
		}

		oldStr := fmt.Sprintf("{%s}", key)
		newStr := vars.Get(key)
		result = strings.ReplaceAll(result, oldStr, newStr)
	}
	return result, nil
}

func parseDate(str string) (string, error) {
	result := str
	dateReg, err := regexp.Compile(`{(Y{4})}|{(M{1,2})}|{(D{1,2})}|{(W{1,4})}`)
	if err != nil {
		return "", err
	}

	for _, match := range dateReg.FindAllStringSubmatch(str, -1) {
		symbol := match[1]
		oldStr := fmt.Sprintf("{%s}", symbol)
		newStr := fmt.Sprintf("%s", parseDateFormat(symbol))
		result = strings.ReplaceAll(result, oldStr, newStr)
	}
	return result, nil
}

var (
	_YYYY = []string{"YYYY", "2006"}
	_MM   = []string{"MM", "01"}
	_M    = []string{"M", "1"}
	_DD   = []string{"DD", "02"}
	_D    = []string{"D", "2"}
	_WWWW = []string{"WWWW", "3"}
	_WWW  = []string{"WWW", "2"}
	_WW   = []string{"WW", "1"}
	_W    = []string{"W", "0"}
)

func parseWeekday(weekday time.Weekday) []string {
	switch weekday {
	case time.Sunday:
		return []string{"日", "Su", "Sun", "Sunday"}
	case time.Monday:
		return []string{"月", "Mo", "Mon", "Monday"}
	case time.Tuesday:
		return []string{"火", "ue", "Tue", "Tuesday"}
	case time.Wednesday:
		return []string{"水", "ne", "Wed", "Wednesday"}
	case time.Thursday:
		return []string{"木", "ur", "Thu", "Thursday"}
	case time.Friday:
		return []string{"金", "Fr", "Fri", "Friday"}
	case time.Saturday:
		return []string{"土", "tu", "Sat", "Saturday"}
	}
}

func parseDateFormat(pattern string) string {
	now := time.Now()

	switch pattern {
	case _YYYY[0]:
		return now.Format(_YYYY[1])
	case _MM[0]:
		return now.Format(_MM[1])
	case _M[0]:
		return now.Format(_M[1])
	case _DD[0]:
		return now.Format(_DD[1])
	case _D[0]:
		return now.Format(_D[1])
	}

	week := parseWeekday(now.Weekday())
	switch pattern {
	case _WWWW[0]:
		return week[3]
	case _WWW[0]:
		return week[2]
	case _WW[0]:
		return week[1]
	case _W[0]:
		return week[0]
	}
	return ""
}
