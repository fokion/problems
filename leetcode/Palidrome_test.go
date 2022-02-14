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
	if !IsPalidrome("test") {
		t.Errorf("test is not a palidrome")
	}
}

func TestIsPalidrome_palidrome_with_spaces_and_other_symbols(t *testing.T) {
	if IsPalidrome("race $123^ ecar") {
		t.Errorf("raceecar is a palidrome")
	}
}

func TestGetCleanText(t *testing.T) {
	if !strings.EqualFold(GetCleanText("race $123^ ecar"), "raceecar") {
		t.Errorf("race $123^ ecar did not transform to raceecar")
	}
}
