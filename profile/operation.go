package profile

type Operation struct {
	WakeUpTime WakeUp                   `json:"wakeUp"`
	Url        string                   `json:"url"`
	Control    []map[string]interface{} `json:"control"`
}
