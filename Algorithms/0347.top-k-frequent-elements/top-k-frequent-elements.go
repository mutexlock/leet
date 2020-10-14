package problem0347

import (
	"container/heap"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	// 统计每个数字出现的次数
	rec := make(map[int]int, len(nums))
	for _, n := range nums {
		rec[n]++
	}
	// 对出现次数进行排序
	counts := make([]int, 0, len(rec))
	for _, c := range rec {
		counts = append(counts, c)
	}
	sort.Ints(counts)
	// min 是 前 k 个高频数字的底线
	min := counts[len(counts)-k]

	// 收集所有　不低于　底线的数字
	for n, c := range rec {
		if c >= min {
			res = append(res, n)
		}
	}

	return res
}


func topKFrequent1(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	q := PriorityQueue{}
	for key, count := range m {
		heap.Push(&q, &Item{key: key, count: count})
	}
	var result []int
	for len(result) < k {
		item := heap.Pop(&q).(*Item)
		result = append(result, item.key)
	}
	return result
}

// Item define
type Item struct {
	key   int
	count int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// 注意：因为golang中的heap是按最小堆组织的，所以count越大，Less()越小，越靠近堆顶.
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push define
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop define
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}