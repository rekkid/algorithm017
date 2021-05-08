//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//
//
// 示例 1：
//
// 输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]
//
//
// 示例 2：
//
// 输入：arr = [0,1,2,1], k = 1
//输出：[0]
//
//
//
// 限制：
//
//
// 0 <= k <= arr.length <= 10000
// 0 <= arr[i] <= 10000
//
// Related Topics 堆 分治算法
// 👍 237 👎 0

package main

import (
	"container/heap"
	"reflect"
	"sort"
	"testing"
)

func TestZuiXiaoDeKgeShuLcof(t *testing.T) {
	var getLeastNumbersTests = []struct {
		arr      []int
		k        int
		expected []int
	}{
		//{[]int{4, 5, 1, 6, 2, 7, 3, 8}, 4, []int{1, 2, 3, 4}},
		//{[]int{3, 2, 1}, 2, []int{1, 2}},
		//{[]int{0, 1, 2, 1}, 1, []int{0}},
		{[]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8, []int{0, 0, 1, 1, 2, 2, 2, 3}},
	}
	for _, tt := range getLeastNumbersTests {
		actual := getLeastNumbers(tt.arr, tt.k)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("getLeastNumbers(%v, %v) = %v; expected %v", tt.arr, tt.k, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Peek() int {
	return (*h)[0]
}

func getLeastNumbers_ref(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	maxHeap := &MaxHeap{}
	for i := 0; i < len(arr); i++ {
		if maxHeap.Len() < k {
			heap.Push(maxHeap, arr[i])
		} else if arr[i] < maxHeap.Peek() {
			heap.Pop(maxHeap)
			heap.Push(maxHeap, arr[i])
		}
	}
	var res []int
	for _, v := range *maxHeap {
		res = append(res, v)
	}
	return res
}

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	maxHeap := MaxHeap(arr[:k])
	heap.Init(&maxHeap)
	for i := k; i < len(arr); i++ {
		if arr[i] < maxHeap[0] {
			heap.Pop(&maxHeap)
			heap.Push(&maxHeap, arr[i])
		}
	}
	result := []int{}
	for _, v := range maxHeap {
		result = append(result, v)
	}
	//sort.Ints(result)
	return result
}

func getLeastNumbers_bad(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	sort.Ints(arr)
	return arr[:k]
}

//leetcode submit region end(Prohibit modification and deletion)
