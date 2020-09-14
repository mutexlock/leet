package problem0231

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%2 == 1 {
			return false
		}
		n /= 2
	}

	return true
}

// 解法一 二进制位操作法
func isPowerOfTwo1(num int) bool {
	return (num > 0 && ((num & (num - 1)) == 0))
}
