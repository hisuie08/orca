package color

import (
	"fmt"
	"regexp"
)

// 色文字用
type StringColor = string

// 色文字適用用
type StringLike interface {
	~string
}

const (
	colorReset StringColor = "\033[0m"
	Red        StringColor = "\033[31m"
	Green      StringColor = "\033[32m"
	Yellow     StringColor = "\033[33m"
	Blue       StringColor = "\033[34m"
	Gray       StringColor = "\033[90m"
)

func Colored[T StringLike](s T, c StringColor) T {
	return T(fmt.Sprintf("%s%s%s", c, s, colorReset))
}

func UnColored[T StringLike](s T) T {
	ansiRegexp := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	ns := ansiRegexp.ReplaceAllString(string(s), "")
	return T(ns)
}

func ColorString[T StringLike](s T, c StringColor, enabled bool) T {
	if !enabled {
		return UnColored(s)
	}
	return Colored(s, c)
}
