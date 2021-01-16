package main

import (
	"math"
	"strconv"
)

type ListNode struct {
	 Val int
	     Next *ListNode
}

func main() {
	// da := reverse(1534236469)
	//	//println(da)
    dt := &ListNode{9,nil}
	l1 := &ListNode{9,new(ListNode)}
	l1.Next = dt
	l2 := &ListNode{ 9,nil}
	pp := addTwoNumbers(l1,l2)
	for {
		if(pp != nil){
			println(pp.Val)
			pp = pp.Next
		}
	}
}

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
		if l2 ==nil || l2.Next == nil {
			l2 = nil
		}else {
			l2 = l2.Next
		}
		addTwoNumbers(temp, l2)
	}
	return l1
}

func reverse(x int) int {
	if math.Pow(2, 31)-1 < float64(x) || math.Pow(-2, 31) > float64(x) {
		return 0
	}
	var result []byte
	isF := false
	if x < 0 {
		isF = true
		result = []byte (strconv.Itoa(-x))
	} else {
		result = []byte(strconv.Itoa(x))
	}
	lenth := len(result)
	neeRes := make([]byte,lenth)

	for index, _ := range result {
		neeRes[lenth-index-1] = result[index]
	}
	valres,_:=strconv.Atoi(string(neeRes[:]))

	if math.Pow(2, 31)-1 < float64(valres) || math.Pow(-2, 31) > float64(valres) {
		return 0
	}
	if isF {
		return -valres
	} else {
		return valres
	}
}
