package profile

import "regexp"

type Profiles []Profile

type Profile struct {
	Name      string    `json:"name"`
	Operation Operation `json:"operation"`
}

type Operation struct {
	WakeUpTime WakeUp              `json:"wakeUp"`
	Url        string              `json:"url"`
	Control    []map[string]string `json:"control"`
}

type WakeUp struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

func (w WakeUp) DateTime() {
	reg := regexp.Compile("")
}
