package profile

import (
	"fmt"
	"regexp"
	"strings"
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
