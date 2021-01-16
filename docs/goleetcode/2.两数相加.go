/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 方法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return l1
	} else {
		val1 := 0
		val2 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		if l2 != nil {
			val2 = l2.Val
		}
		sum := val1 + val2
		var temp *ListNode
		temp = l1.Next

		if sum >= 10 {
			l1.Val = sum % 10
			if temp != nil {
				temp.Val = temp.Val + 1
			} else {
				temp.Val = 1
			}
		} else {
			l1.Val = sum
		}
		l1.Next = temp
		if l2 == nil || l2.Next == nil {
			l2 = nil
		} else {
			l2 = l2.Next
		}
		addTwoNumbers(temp, l2)
	}
	return l1
}

// @lc code=end
