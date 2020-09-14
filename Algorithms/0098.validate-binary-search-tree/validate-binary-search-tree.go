package problem0098

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
	"math"
)

type TreeNode = kit.TreeNode

func isValidBST(root *TreeNode) bool {
	// Go int 类型的最小值与最大值
	MIN, MAX := -1<<63, 1<<63-1

	return recur(MIN, MAX, root)
}

// 以递归的方式，检查 root.Val 是否在 (min, max) 范围内。
func recur(min, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}

	return min < root.Val && root.Val < max &&
		recur(min, root.Val, root.Left) &&
		recur(root.Val, max, root.Right)
}

// 解法一，直接按照定义比较大小，比 root 节点小的都在左边，比 root 节点大的都在右边
func isValidBST1(root *TreeNode) bool {
	return isValidbst(root, math.Inf(-1), math.Inf(1))
}
func isValidbst(root *TreeNode, min, max float64) bool {
	if root == nil {
		return true
	}
	v := float64(root.Val)
	return v < max && v > min && isValidbst(root.Left, min, v) && isValidbst(root.Right, v, max)
}

// 解法二，把 BST 按照左中右的顺序输出到数组中，如果是 BST，则数组中的数字是从小到大有序的，如果出现逆序就不是 BST
func isValidBST2(root *TreeNode) bool {
	arr := []int{}
	inOrder(root, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inOrder(root.Right, arr)
}
