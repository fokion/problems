package leetcode

import (
	"fmt"
	"regexp"
	"strings"
)

func IsPalidrome(text string) bool {
	text = GetCleanText(text)
	if len(text) <= 1 {
		return true
	}
	//divide text in half
	index := len(text) / 2
	firstHalf := text[:index]
	secondHalf := text[index:]
	if len(text)%2 == 1 {
		firstHalf = text[:index]
		secondHalf = text[index+1:]
	}
	length := len(firstHalf)
	fmt.Println(firstHalf, secondHalf, length)
	for i := 0; i < length; i++ {
		fmt.Println(string(firstHalf[i]), string(secondHalf[(length-i-1)]))
		if !strings.EqualFold(string(firstHalf[i]), string(secondHalf[(length-i-1)])) {
			return false
		}
	}

	return true
}

func GetCleanText(text string) string {

	expr := regexp.MustCompile(`(\s+|([^a-z0-9]+))`)
	return expr.ReplaceAllString(strings.ToLower(text), "")
}
