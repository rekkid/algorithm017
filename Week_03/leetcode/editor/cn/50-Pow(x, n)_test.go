//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。
//
//
//
// 示例 1：
//
//
//输入：x = 2.00000, n = 10
//输出：1024.00000
//
//
// 示例 2：
//
//
//输入：x = 2.10000, n = 3
//输出：9.26100
//
//
// 示例 3：
//
//
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2-2 = 1/22 = 1/4 = 0.25
//
//
//
//
// 提示：
//
//
// -100.0 < x < 100.0
// -231 <= n <= 231-1
// -104 <= xn <= 104
//
// Related Topics 数学 二分查找
// 👍 643 👎 0

package main

import (
	"reflect"
	"testing"
)

func TestPowxN(t *testing.T) {
	var myPowTests = []struct {
		x        float64
		n        int
		expected float64
	}{
		{2.00000, 10, 1024.00000},
		{2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
	}
	for _, tt := range myPowTests {
		actual := myPow(tt.x, tt.n)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("myPow(%v, %v) = %v; expected %v", tt.x, tt.n, actual, tt.expected)
		}
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := myPowHelper(x, n)
	return res
}

func myPowHelper(x float64, level int) float64 {
	// !!陷阱一定要注意
	if level == 0 {
		return 1
	}
	if level == 1 {
		return x
	}
	tmp := myPowHelper(x, level/2)
	if level%2 == 0 {
		return tmp * tmp
	} else {
		return tmp * tmp * x
	}
}

//leetcode submit region end(Prohibit modification and deletion)
