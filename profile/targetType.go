package profile

type TargetType string

const (
	IsCSSSelector = TargetType("CSSSelector")
	IsXPath       = TargetType("XPath")
)
