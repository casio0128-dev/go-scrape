package profile

import (
	"go-scrape/common"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Profiles []Profile

type Profile struct {
	Name      string    `json:"name"`
	Variable  Variable  `json:"var"`
	Operation Operation `json:"operation"`
}

type Variable map[string]string

func (v *Variable) get(key string) string {
	if val, ok := (*v)[key]; ok {
		return val
	}
	return ""
}

func (v *Variable) set(key, value string) {
	(*v)[key] = value
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

func (w WakeUp) DateTime() (time.Time, error) {
	now := time.Now()

	var (
		year   = now.Year()
		month  = int(now.Month())
		day    = now.Day()
		hour   = now.Hour()
		minute = now.Minute()

		err error
	)

	if w.isActualDateFormat() {
		dateSep := strings.Split(w.Date, common.DATE_SEPARATER)
		if year, err = strconv.Atoi(dateSep[0]); err != nil {
			return time.Time{}, err
		}
		if month, err = strconv.Atoi(dateSep[1]); err != nil {
			return time.Time{}, err
		}
		if day, err = strconv.Atoi(dateSep[2]); err != nil {
			return time.Time{}, err
		}
	}

	if w.isActualDateFormat() {
		timeSep := strings.Split(w.Time, common.TIME_SEPARATER)
		if hour, err = strconv.Atoi(timeSep[0]); err != nil {
			return time.Time{}, err
		}
		if minute, err = strconv.Atoi(timeSep[1]); err != nil {
			return time.Time{}, err
		}
	}

	loc, err := time.LoadLocation("Local")
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, loc), nil
}

func (w WakeUp) isActualDateFormat() bool {
	dateRegex, err := regexp.Compile(common.DATE_REGEXP)
	if err != nil {
		panic(err)
	}
	return dateRegex.MatchString(w.Date)
}

func (w WakeUp) isActualTimeFormat() bool {
	timeRegex, err := regexp.Compile(common.TIME_REGEXP)
	if err != nil {
		panic(err)
	}
	return timeRegex.MatchString(w.Time)
}
