package problem0141

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow == fast
}


func permute(nums []int) [][]int {

	res := make([]int,0)
	var result [][]int
	backtrack(nums,res, result)
	return result
}

func backtrack(nums []int, res []int, result [][]int){
	if len(res) == len(nums){
		result = append(result,res)
		return
	}
	for _,n := range nums{
		if contains(res,n){
			continue
		}
		res = append(res,n)
		backtrack(nums,res,result)
		res = res[:len(res)-1]
	}
}

func contains(nums []int,k int)bool  {
	for _,n := range nums{
		if n  == k {
			return  true
		}
	}
	return  false
}
//List<List<Integer>> res = new LinkedList<>();
//
///* 主函数，输入一组不重复的数字，返回它们的全排列 */
//List<List<Integer>> permute(int[] nums) {
//// 记录「路径」
//LinkedList<Integer> track = new LinkedList<>();
//backtrack(nums, track);
//return res;
//}
//
//// 路径：记录在 track 中
//// 选择列表：nums 中不存在于 track 的那些元素
//// 结束条件：nums 中的元素全都在 track 中出现
//void backtrack(int[] nums, LinkedList<Integer> track) {
//// 触发结束条件
//if (track.size() == nums.length) {
//res.add(new LinkedList(track));
//return;
//}
//
//for (int i = 0; i < nums.length; i++) {
//// 排除不合法的选择
//if (track.contains(nums[i]))
//continue;
//// 做选择
//track.add(nums[i]);
//// 进入下一层决策树
//backtrack(nums, track);
//// 取消选择
//track.removeLast();
//}
//}