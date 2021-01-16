package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	result := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		val := target - nums[i]
		sign := false
		for index, value := range nums {
			if value == val && index != i {
				sign = true
				result[0] = i
				result[1] = index
				break
			}
		}
		if sign {
			break
		}
	}
	return result
}

// @lc code=end
