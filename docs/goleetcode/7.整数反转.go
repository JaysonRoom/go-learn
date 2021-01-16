/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	if math.Pow(2, 31)-1 < float64(x) && math.Pow(-2, -31) > float64(x) {
		return 0
	}
	var result []byte
	isF := false
	if x < 0 {
		isF = true
		result = []byte(strconv.Itoa(-x))
	} else {
		result = []byte(strconv.Itoa(x))
	}
	lenth := len(result)
	neeRes := make([]byte, lenth)

	for index, _ := range result {
		neeRes[lenth-index-1] = result[index]
	}
	valres, _ := strconv.Atoi(string(neeRes[:]))

	if math.Pow(2, 31)-1 < float64(valres) || math.Pow(-2, 31) > float64(valres) {
		return 0
	}
	if isF {
		return -valres
	} else {
		return valres
	}
}

// @lc code=end

