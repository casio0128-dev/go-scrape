package profile

type Profiles []Profile

type Profile struct {
	Name      string    `json:"name"`
	Operation Operation `json:"operation"`
}

type Operation struct {
	WakeUpTime string              `json:"wakeUp"`
	Url        string              `json:"url"`
	Control    []map[string]string `json:"control"`
}
