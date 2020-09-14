package problem0102

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的 level
		if  level >= len(res)   {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}

//better
// 解法一 BFS
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	curNum, nextLevelNum, res, tmp := 1, 0, [][]int{}, []int{}
	for len(queue) != 0 {
		if curNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			curNum--
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}
		if curNum == 0 {
			res = append(res, tmp)
			curNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
		}
	}
	return res
}

// 解法二 DFS
func levelOrder2(root *TreeNode) [][]int {
	levels := [][]int{}
	dfsLevel(root, -1, &levels)
	return levels
}

func dfsLevel(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}
	currLevel := level + 1
	for len(*res) <= currLevel {
		*res = append(*res, []int{})
	}
	(*res)[currLevel] = append((*res)[currLevel], node.Val)
	dfsLevel(node.Left, currLevel, res)
	dfsLevel(node.Right, currLevel, res)
}