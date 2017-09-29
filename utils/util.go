package utils

type LimitOffset struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}
