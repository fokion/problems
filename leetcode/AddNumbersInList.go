package leetcode

import (
	"fmt"
	"github.com/fokion/problems/leetcode/models"
	"math"
	"reflect"
	"strconv"
)

func AddTwoNumbers(firstList *models.ListNode, secondList *models.ListNode) *models.ListNode {
	return add(firstList, secondList, false)
}
func finishList(l1 *models.ListNode, carry bool) *models.ListNode {
	var node *models.ListNode
	sum := l1.Val
	if carry {
		sum++
	}
	node = &models.ListNode{Val: sum % 10}
	if !reflect.ValueOf(l1.Next).IsNil() {
		node.Next = finishList(l1.Next, sum >= 10)
	}
	return node
}
func add(l1 *models.ListNode, l2 *models.ListNode, carry bool) *models.ListNode {
	var node *models.ListNode
	if !reflect.ValueOf(l1).IsNil() || !reflect.ValueOf(l2).IsNil() {
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		sum := val1 + val2
		if carry {
			sum++
		}
		node = &models.ListNode{Val: sum % 10}

		if reflect.ValueOf(l1.Next).IsNil() {
			if reflect.ValueOf(l2.Next).IsNil() {
				node.Next = finishList(&models.ListNode{Val: 0}, sum >= 10)
			} else {
				node.Next = finishList(l2.Next, sum >= 10)
			}
		} else if reflect.ValueOf(l2.Next).IsNil() {
			node.Next = finishList(l1.Next, sum >= 10)
		} else {
			node.Next = add(l1.Next, l2.Next, sum >= 10)
		}
	} else if carry {
		node = &models.ListNode{Val: 1}
	}
	return node
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
func convertToNumber(list *models.ListNode, times int, sum int) int {
	num := int(float64(list.Val) * math.Pow(10, float64(times)))
	sum += num
	if list.Next != nil {
		return convertToNumber(list.Next, times+1, sum)
	}
	return sum
}
