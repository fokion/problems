package leetcode

import (
	"strings"
	"testing"
)

func TestIsPalidrome_anna(t *testing.T) {
	if !IsPalidrome("anna") {
		t.Errorf("anna is a palidrome")
	}
}

func TestIsPalidrome_test(t *testing.T) {
	if IsPalidrome("test") {
		t.Errorf("test is not a palidrome")
	}
}

func TestIsPalidrome_palidrome_with_spaces_and_other_symbols(t *testing.T) {
	if !IsPalidrome("race $^ ecar") {
		t.Errorf("raceecar is a palidrome")
	}
}

func TestIsPalidrome(t *testing.T) {
	if !IsPalidrome("A man, a plan, a canal: Panama") {
		t.Errorf("A man, a plan, a canal: Panama -> amanaplanacanalpanama which is a palidrome")
	}
}
func TestGetCleanText_keep_numbers(t *testing.T) {
	if !strings.EqualFold(GetCleanText("1234 $!"), "1234") {
		t.Errorf("1234 $! did not transform to 1234 , %s", GetCleanText("1234 $!"))
	}
}
func TestGetCleanText(t *testing.T) {
	if !strings.EqualFold(GetCleanText("race $^ ecar"), "raceecar") {
		t.Errorf("race $^ ecar did not transform to raceecar")
	}
}
