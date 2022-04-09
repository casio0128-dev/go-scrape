package profile

type Variable map[string]string

func (vars *Variable) IsExists(key string) bool {
	_, ok := (*vars)[key]
	return ok
}

func (vars *Variable) Get(key string) string {
	if val, ok := (*vars)[key]; ok {
		return val
	}
	return ""
}

func (vars *Variable) Set(key, value string) {
	(*vars)[key] = value
}
