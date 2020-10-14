package problem0239

// 参看 https://leetcode.com/problems/sliding-window-maximum/discuss/65881/O(n)-solution-in-Java-with-two-simple-pass-in-the-array
func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	if k <= 1 {
		return nums
	}

	g := k - 1 // 比参考文章的分组少一个，可以减少 max 函数的调用，理论上可以加速。

	left := make([]int, size)
	for i := 0; i < size; i++ {
		if i%g == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(nums[i], left[i-1])
		}
	}

	right := make([]int, size)
	// size-1 很可能不是那组的最后一个，需要单独列出
	right[size-1] = nums[size-1]
	for j := size - 2; j >= 0; j-- {
		if (j+1)%g == 0 {
			right[j] = nums[j]
		} else {
			right[j] = max(nums[j], right[j+1])
		}
	}

	res := make([]int, size-k+1)
	for i := 0; i <= size-k; i++ {
		// right[i] 中保存了 nums[i:g*(i/g+1)] 中的最大值
		// left[i+k-1] 中保存了 nums[g*(i/g+1):i+k] 中的最大值
		res[i] = max(right[i], left[i+k-1])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	window := make([]int, 0, k) // store the index of nums
	result := make([]int, 0, len(nums)-k+1)
	for i, v := range nums { // if the left-most index is out of window, remove it
		if i >= k && window[0] <= i-k {
			window = window[1:len(window)]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < v { // maintain window
			window = window[0 : len(window)-1]
		}
		window = append(window, i) // store the index of nums
		if i >= k-1 {
			result = append(result, nums[window[0]]) // the left-most is the index of max value in nums
		}
	}
	return result
}