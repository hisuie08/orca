package formatter

import (
	"fmt"
	"orca/presenter/formatter/color"
	"strings"
	"unicode/utf8"
)

type StringLike interface {
	~string
}

// Shorten returns a StringLike T truncated to the maximum length of <width>.
// <width> as total length including <suffix>, if it not default
func Shorten[T StringLike](s T, width int, suffix string) T {
	length := min(width-len(suffix), VisibleLen(s))
	return T(fmt.Sprintf("%s%s", s[:length], suffix))
}

func VisibleLen[T StringLike](s T) int {
	us := string(color.UnColored(s))
	return utf8.RuneCountInString(us)
}

// PadWidh returns <s> with padding if it's shorter so that its length is <width>.
func PadWidth[T StringLike](s T, width int) T {
	pad := max(width-VisibleLen(s), 0)
	return T(fmt.Sprintf("%s%s", s, strings.Repeat(" ", pad)))
}
