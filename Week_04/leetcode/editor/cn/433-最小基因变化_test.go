//一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
//
// 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
//
// 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
//
// 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
//
// 现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变
//化次数。如果无法实现目标变化，请返回 -1。
//
// 注意：
//hg
//
// 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
// 如果一个起始基因序列需要多次变化，那么它每一次变化之后的基因序列都必须是合法的。
// 假定起始基因序列与目标基因序列是不一样的。
//
//
//
//
// 示例 1：
//
//
//start: "AACCGGTT"
//end:   "AACCGGTA"
//bank: ["AACCGGTA"]
//
//返回值: 1
//
//
// 示例 2：
//
//
//start: "AACCGGTT"
//end:   "AAACGGTA"
//bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
//
//返回值: 2
//
//
// 示例 3：
//
//
//start: "AAAAACCC"
//end:   "AACCCCCC"
//bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
//
//返回值: 3
//
// 👍 72 👎 0

package main

import (
	"reflect"
	"testing"
)

func TestMinimumGeneticMutation(t *testing.T) {
	var minMutationTests = []struct {
		start    string
		end      string
		bank     []string
		expected int
	}{
		{"AAAACCCC", "CCCCCCCC", []string{"AAAACCCA", "AAACCCCA", "AACCCCCA", "AACCCCCC", "ACCCCCCC", "CCCCCCCC", "AAACCCCC", "AACCCCCC"}, 1},
		//{"AA", "CC", []string{"AG", "AC", "GA", "CA", "CC"}, 2},
		//{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
		//{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
		//{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, 3},
		//{"AACCGGTT", "AACCGCTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
		//{"AACCTTGG", "AATTCCGG", []string{"AATTCCGG", "AACCTGGG", "AACCCCGG", "AACCTACC"}, -1},
		//{"AAAAAAAA", "CCCCCCCC", []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC", "AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA", "CCCCCCCC"}, 8},
	}
	for _, tt := range minMutationTests {
		actual := minMutation(tt.start, tt.end, tt.bank)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("minMutation(%v, %v, %v) = %v; expected %v", tt.start, tt.end, tt.bank, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)

var changeMap = map[string]string{
	"A": "CGT",
	"C": "AGT",
	"G": "ACT",
	"T": "ACG",
}

func isFind(word string, bank []string) (int, bool) {
	for i, v := range bank {
		if word == v {
			return i, true
		}
	}
	return -1, false
}

// 使用BFS进行推导
func minMutation_BFS(start string, end string, bank []string) int {
	if _, ok := isFind(end, bank); !ok {
		return -1
	}
	step := 0
	queue := []string{start}
	visited := make([]bool, len(bank))
	for len(queue) != 0 {
		var tmp []string
		for _, current := range queue {
			if current == end {
				//fmt.Println("found")
				return step
			}

			for i, char := range current {
				for _, value := range changeMap[string(char)] {
					new := current[:i] + string(value) + current[i+1:]
					//fmt.Println("    new: ", new)
					if pos, found := isFind(new, bank); found && !visited[pos] {
						//fmt.Println("current: ", current)
						//fmt.Println("    new: ", new)
						visited[pos] = true
						tmp = append(tmp, new)
					}
				}
			}
		}
		step++
		queue = tmp
	}
	return -1
}

// 使用DFS递归进行遍历
func minMutation(start string, end string, bank []string) int {

	visited := make(map[string]bool)
	var DFS func(current string, res []string, step int) int
	DFS = func(current string, res []string, step int) int {
		if visited[current] == true {
			return step
		}
		visited[current] = true
		if current == end {
			//fmt.Println(res)
			//fmt.Println("found")
			if step == -1 {
				step = len(res) - 1
			} else if step > len(res)-1 {
				step = len(res) - 1
			}
			return step
		}
		for i, char := range current {
			for _, value := range changeMap[string(char)] {
				new := current[:i] + string(value) + current[i+1:]
				if _, ok := isFind(new, bank); ok && !visited[new] {
					//tmp := make([]string, len(res))
					//copy(tmp, res)
					//tmp = append(tmp, new)
					res = append(res, new)
					step = DFS(new, res, step)
					res = res[:len(res)-1]
					visited[new] = false
				}
			}
		}
		return step
	}
	step := DFS(start, []string{start}, -1)
	return step
}

//leetcode submit region end(Prohibit modification and deletion)
