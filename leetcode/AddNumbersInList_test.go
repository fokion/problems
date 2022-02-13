package leetcode

import (
	"fmt"
	"github.com/fokion/problems/leetcode/models"
	"testing"
)

func TestConverterToInteger(t *testing.T) {
	//list [2,4,3] must be 342 as a number
	list := models.ListNode{Val: 2}
	var pointerToList *models.ListNode = &list
	list.Next = &models.ListNode{Val: 4, Next: &models.ListNode{Val: 3}}

	got := convertToNumber(pointerToList, 0, 0)
	if 342 != got {
		t.Errorf("The converter is not working properly as I was expecting 342 for [2,3,4] list and got : %d", got)
	}
}
func TestConvertingIntegerToList(t *testing.T) {
	//integer 342 will make a list of [2,4,3]
	got := convertNumberToList(342)
	fistDigit := got.Val
	fmt.Println(got)
	if got.Next == nil {
		t.Errorf("does not have second digit")

	}
	secondDigit := got.Next.Val
	thirdDigit := got.Next.Next.Val

	if thirdDigit != 3 && secondDigit != 4 && fistDigit != 2 {
		t.Errorf("the sequence is not correct %d,%d,%d", fistDigit, secondDigit, thirdDigit)
	}
}
func TestConvertingNumbersBackandForth(t *testing.T) {

	listOne := convertNumberToList(342)
	if 342 != convertToNumber(listOne, 0, 0) {
		t.Errorf("did not convert that list correctly expected 342 , got :  %d", convertToNumber(listOne, 0, 0))
	}
}

func TestSumOfNumbersWithSameLength(t *testing.T) {
	//list one 2,4,3
	listOne := convertNumberToList(342)
	//list two 5,6,4
	listTwo := convertNumberToList(465)
	//expected 7,0,8
	got := AddTwoNumbers(listOne, listTwo)

	gotNumber := convertToNumber(got, 0, 0)
	if gotNumber != 807 {
		t.Errorf("Expected 807 and got %d", gotNumber)
	}
}
