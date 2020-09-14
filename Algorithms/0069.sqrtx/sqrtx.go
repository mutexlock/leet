package problem0069

func mySqrt(x int) int {
	res := x
	// 牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}


// 解法一 二分
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}
	left, right, res := 1, x, 0
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid < x/mid {
			left = mid + 1
			res = mid
		} else if mid == x/mid {
			return mid
		} else {
			right = mid - 1
		}
	}
	return res
}

// 解法二 牛顿迭代法 https://en.wikipedia.org/wiki/Integer_square_root
func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}