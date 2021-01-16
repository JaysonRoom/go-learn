package main

import (
	"math"
	"strconv"
)

// ListNode 结点
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//	l1 := ListNode{2, 3, 4}

}

// 方法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val1 := listNodetoInt(l1)
	val2 := listNodetoInt(l2)
	val3 := val1 + val2
	return intToListNode(val3)
}

func listNodetoInt(ll *ListNode) int {
	i := 0
	sum := 0
	for ll.Next != nil {
		sum += ll.Val * int(math.Pow10(i))
		ll = ll.Next
		i = i + 1
	}
	return sum
}

func intToListNode(val int) *ListNode {
	var result *ListNode
	str := []byte(strconv.Itoa(val))
	for index := range str {
		result.Val = int(str[index])
		temp := new(ListNode)
		result.Next = temp
		result = temp
	}
	return result
}
