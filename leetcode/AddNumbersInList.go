package leetcode

import (
	"fmt"
	"github.com/fokion/problems/leetcode/models"
	"math"
	"strconv"
)

func AddTwoNumbers(firstList *models.ListNode, secondList *models.ListNode) *models.ListNode {

	sum := convertToNumber(firstList, 0, 0) + convertToNumber(secondList, 0, 0)

	return convertNumberToList(sum)
}

func convertToNumber(list *models.ListNode, times int, sum int) int {
	num := int(float64(list.Val) * math.Pow(10, float64(times)))
	sum += num
	if list.Next != nil {
		return convertToNumber(list.Next, times+1, sum)
	}
	return sum
}
func convertNumberToList(number int) *models.ListNode {
	strArr := []rune(strconv.Itoa(number))
	list := models.ListNode{}
	var pointerToList *models.ListNode = &list
	var next *models.ListNode = pointerToList
	for i := len(strArr) - 1; i >= 0; i-- {
		elem := next
		char := string(strArr[i])
		fmt.Println("converting : ", char)
		var num int
		num, _ = strconv.Atoi(char)
		elem.Val = num
		if i-1 >= 0 {
			elem.Next = &models.ListNode{}
		}
		next = elem.Next
	}

	return pointerToList
}
