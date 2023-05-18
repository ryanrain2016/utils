package utils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		name   string
		seq    []int
		r      int
		result [][]int
	}{
		{
			name:   "test case 1",
			seq:    []int{1, 2, 3},
			r:      2,
			result: [][]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			name:   "test case 2",
			seq:    []int{1, 2, 3, 4},
			r:      3,
			result: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([][]int, 0)
			for comb := range Combinations(tt.seq, tt.r) {
				result = append(result, comb)
			}
			if !reflect.DeepEqual(result, tt.result) {
				t.Errorf("Combinations() = %v, want %v", result, tt.result)
			}
		})
	}
}

func TestCombinationsWithReplacement(t *testing.T) {
	tests := []struct {
		name   string
		seq    []int
		r      int
		result [][]int
	}{
		{
			name:   "test case 2",
			seq:    []int{1, 2},
			r:      3,
			result: [][]int{{1, 1, 1}, {1, 1, 2}, {1, 2, 2}, {2, 2, 2}},
		},
		{
			name:   "test case 1",
			seq:    []int{1, 2, 3},
			r:      2,
			result: [][]int{{1, 1}, {1, 2}, {1, 3}, {2, 2}, {2, 3}, {3, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([][]int, 0)
			for comb := range CombinationsWithReplacement(tt.seq, tt.r) {
				result = append(result, comb)
			}
			// if !reflect.DeepEqual(result, tt.result) {
			// 	t.Errorf("CombinationsWithReplacement() = %v, want %v", result, tt.result)
			// }
			assert.ElementsMatch(t, result, tt.result)
		})
	}
}

// BenchmarkCombinationsWithReplacement benchmarks the performance of the CombinationsWithReplacement function.
// The function generates all combinations with replacement of length r for a sequence of n elements.
func BenchmarkCombinationsWithReplacement(b *testing.B) {
	seq := make([]int, 20)
	for i := 0; i < 20; i++ {
		seq[i] = i
	}

	// Benchmark the performance of the CombinationsWithReplacement function for different values of r
	for _, r := range []int{5} {
		b.Run(fmt.Sprintf("r=%d", r), func(b *testing.B) {
			// Measure the time it takes to generate all combinations
			for i := 0; i < b.N; i++ {
				count := 0
				for range CombinationsWithReplacement(seq, r) {
					count++
				}
			}
		})
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		name   string
		start  int
		step   int
		length int
	}{
		{"start=0, step=1", 0, 1, 10},
		{"start=1, step=2", 1, 2, 5},
		{"start=-3, step=3", -3, 3, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := Count(tt.start, tt.step)
			defer func() { // 确保通道被关闭
				close(count)
				for range count {
				}
			}()

			for i := 0; i < tt.length; i++ {
				val := <-count
				expected := tt.start + i*tt.step
				if val != expected {
					t.Errorf("Count() = %v, want %v", val, expected)
				}
			}
		})
	}
}

func TestPermutations(t *testing.T) {
	// 测试用例1：生成一个整数切片的所有长度为 1 的排列
	expected1 := [][]int{{1}, {2}, {3}}
	actual1 := [][]int{}
	for perm := range Permutations([]int{1, 2, 3}, 1) {
		actual1 = append(actual1, perm)
	}
	assert.ElementsMatch(t, expected1, actual1)

	// 测试用例2：生成一个整数切片的所有长度为 2 的排列
	expected2 := [][]int{{1, 2}, {1, 3}, {2, 1}, {2, 3}, {3, 1}, {3, 2}}
	actual2 := [][]int{}
	for perm := range Permutations([]int{1, 2, 3}, 2) {
		actual2 = append(actual2, perm)
	}
	assert.ElementsMatch(t, expected2, actual2)

	// 测试用例3：生成一个整数切片的所有长度为 3 的排列
	expected3 := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	actual3 := [][]int{}
	for perm := range Permutations([]int{1, 2, 3}, 3) {
		actual3 = append(actual3, perm)
	}
	assert.ElementsMatch(t, expected3, actual3)

	// 测试用例4：生成一个空切片的所有长度为 0 的排列
	expected4 := [][]int{{}}
	actual4 := [][]int{}
	for perm := range Permutations([]int{}, 0) {
		actual4 = append(actual4, perm)
	}
	assert.ElementsMatch(t, expected4, actual4)

	// 测试用例5：生成一个空切片的所有长度为 1 的排列
	expected5 := [][]int{}
	actual5 := [][]int{}
	for perm := range Permutations([]int{}, 1) {
		actual5 = append(actual5, perm)
	}
	assert.ElementsMatch(t, expected5, actual5)

	// 测试用例6：生成一个长度为 1 的切片的所有长度为 1 的排列
	expected6 := [][]int{{1}}
	actual6 := [][]int{}
	for perm := range Permutations([]int{1}, 1) {
		actual6 = append(actual6, perm)
	}
	assert.ElementsMatch(t, expected6, actual6)
}

func TestProduct(t *testing.T) {
	// 测试用例1：生成两个整数切片的笛卡尔积
	expected1 := [][]int{
		{1, 2, 1, 2},
		{1, 2, 1, 3},
		{1, 2, 2, 2},
		{1, 2, 2, 3},
		{1, 3, 1, 2},
		{1, 3, 1, 3},
		{1, 3, 2, 2},
		{1, 3, 2, 3},
		{2, 2, 1, 2},
		{2, 2, 1, 3},
		{2, 2, 2, 2},
		{2, 2, 2, 3},
		{2, 3, 1, 2},
		{2, 3, 1, 3},
		{2, 3, 2, 2},
		{2, 3, 2, 3},
	}
	actual1 := [][]int{}
	for prod := range Product([][]int{{1, 2}, {2, 3}}, 2) {
		actual1 = append(actual1, prod)
	}
	assert.ElementsMatch(t, expected1, actual1)

	// 测试用例2：生成两个字符串切片的笛卡尔积
	expected2 := [][]string{
		{"a", "x"},
		{"a", "y"},
		{"b", "x"},
		{"b", "y"},
	}
	actual2 := [][]string{}
	for prod := range Product([][]string{{"a", "b"}, {"x", "y"}}, 1) {
		actual2 = append(actual2, prod)
	}
	assert.ElementsMatch(t, expected2, actual2)

	// 测试用例3：生成一个整数切片和一个空切片的笛卡尔积
	expected3 := [][]int{}
	actual3 := [][]int{}
	for prod := range Product([][]int{{1, 2, 3}, {}}, 1) {
		actual3 = append(actual3, prod)
	}
	assert.ElementsMatch(t, expected3, actual3)

	// 测试用例4：生成一个空切片和一个整数切片的笛卡尔积
	expected4 := [][]int{}
	actual4 := [][]int{}
	for prod := range Product([][]int{{}, {1, 2, 3}}, 1) {
		actual4 = append(actual4, prod)
	}
	assert.ElementsMatch(t, expected4, actual4)

	// 测试用例5：生成一个整数切片的笛卡尔积，其中重复次数为 1
	expected5 := [][]int{
		{1},
		{2},
		{3},
	}
	actual5 := [][]int{}
	for prod := range Product([][]int{{1, 2, 3}}, 1) {
		actual5 = append(actual5, prod)
	}
	assert.ElementsMatch(t, expected5, actual5)

	// 测试用例6：生成一个整数切片的笛卡尔积，其中重复次数为 0
	expected6 := [][]int{{}}
	actual6 := [][]int{}
	for prod := range Product([][]int{{1, 2, 3}}, 0) {
		actual6 = append(actual6, prod)
	}
	assert.ElementsMatch(t, expected6, actual6)
}

func BenchmarkProduct(b *testing.B) {
	slice := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	repeat := 3
	for i := 0; i < b.N; i++ {
		out := Product(slice, repeat)
		for range out {
		}
	}
}
