package leetcode

import (
	"regexp"
	"strings"
)

func IsPalidrome(text string) bool {
	text = GetCleanText(text)
	return false
}

func GetCleanText(text string) string {

	expr := regexp.MustCompile(`(\d+|\s+|\W)`)
	return expr.ReplaceAllString(strings.ToLower(text), "")
}
