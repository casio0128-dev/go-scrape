package profile

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Profiles []Profile

type Profile struct {
	Name       string     `json:"name"`
	TargetType TargetType `json:"targetIs"`
	Variable   Variable   `json:"var"`
	Operation  Operation  `json:"operation"`
}

func Parse(str string, vars Variable) (string, error) {
	var result string
	result, err := parseDateTime(str)
	if err != nil {
		return "", err
	}

	reg, err := regexp.Compile(`{\D([a-z|A-Z|\d]*)}`)
	if err != nil {
		return "", err
	}
	matches := reg.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		oldStr := match[0]
		key := oldStr[1 : len(oldStr)-1]
		if !vars.IsExists(key) {
			continue
		}

		newStr := vars.Get(key)
		result = strings.ReplaceAll(result, oldStr, newStr)
	}
	return result, nil
}

func parseDateTime(str string) (string, error) {
	result := str
	dateReg, err := regexp.Compile(`{(Y{4})}|{(Y{2})}|{(M{1,2})}|{(D{1,2})}|{(W{1,4})}|{(H{2})}|{(h{1,2})}|{(m{1,2})}|{(s{1,2})}|{(AMPM)}|{(ampm)}`)
	if err != nil {
		return "", err
	}

	for _, match := range dateReg.FindAllStringSubmatch(str, -1) {
		oldStr := match[0]
		symbol := oldStr[1 : len(oldStr)-1]
		newStr := fmt.Sprintf("%s", parseDateFormat(symbol))
		result = strings.ReplaceAll(result, oldStr, newStr)
	}
	return result, nil
}

var (
	_YYYY  = []string{"YYYY", "2006"}
	_YY    = []string{"YY", "06"}
	_MM    = []string{"MM", "01"}
	_M     = []string{"M", "1"}
	_DD    = []string{"DD", "02"}
	_D     = []string{"D", "2"}
	_HH    = []string{"HH", "15"}
	_hh    = []string{"hh", "03"}
	_h     = []string{"h", "3"}
	_mm    = []string{"mm", "04"}
	_m     = []string{"m", "4"}
	_ss    = []string{"ss", "05"}
	_s     = []string{"s", "5"}
	_AM_PM = []string{"AMPM", "PM"}
	_am_pm = []string{"ampm", "pm"}
)

func getWeekday(weekPattern string) string {
	index := strings.Count(weekPattern, "W") - 1
	if index < 0 || index > 3 {
		return ""
	}
	switch time.Now().Weekday() {
	case time.Sunday:
		return []string{"日", "Su", "Sun", "Sunday"}[index]
	case time.Monday:
		return []string{"月", "Mo", "Mon", "Monday"}[index]
	case time.Tuesday:
		return []string{"火", "Tu", "Tue", "Tuesday"}[index]
	case time.Wednesday:
		return []string{"水", "We", "Wed", "Wednesday"}[index]
	case time.Thursday:
		return []string{"木", "Th", "Thu", "Thursday"}[index]
	case time.Friday:
		return []string{"金", "Fr", "Fri", "Friday"}[index]
	case time.Saturday:
		return []string{"土", "Sa", "Sat", "Saturday"}[index]
	}
	return ""
}

func parseDateFormat(pattern string) string {
	now := time.Now()

	switch pattern {
	case _YYYY[0]:
		return now.Format(_YYYY[1])
	case _YY[0]:
		return now.Format(_YY[1])
	case _MM[0]:
		return now.Format(_MM[1])
	case _M[0]:
		return now.Format(_M[1])
	case _DD[0]:
		return now.Format(_DD[1])
	case _D[0]:
		return now.Format(_D[1])
	case _HH[0]:
		return now.Format(_HH[1])
	case _hh[0]:
		return now.Format(_hh[1])
	case _h[0]:
		return now.Format(_h[1])
	case _mm[0]:
		return now.Format(_mm[1])
	case _m[0]:
		return now.Format(_m[1])
	case _ss[0]:
		return now.Format(_ss[1])
	case _s[0]:
		return now.Format(_s[1])
	case _AM_PM[0]:
		return now.Format(_AM_PM[1])
	case _am_pm[0]:
		return now.Format(_am_pm[1])
	}
	return getWeekday(pattern)
}
