package task1

import (
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	type test struct {
		nums []int
		want int
	}
	tests := map[string]test{
		"t1": {nums: []int{2, 2, 1}, want: 1},
		"t2": {nums: []int{4, 1, 2, 1, 2}, want: 4},
		"t3": {nums: []int{1}, want: 1},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := SingleNumber(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %v, got: %v, want: %v", name, got, tc.want)
			}
		})
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleNumber([]int{2, 2, 1})
	}
}

func TestIsPalindrome(t *testing.T) {
	type test struct {
		x    int
		want bool
	}

	tests := map[string]test{
		"t1": {x: 121, want: true},
		"t2": {x: -121, want: false},
		"t3": {x: 10, want: false},
	}
	for name, tc := range tests {
		got := IsPalindrome(tc.x)
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %v, got: %v, want: %v", name, got, tc.want)
			}
		})
	}
}

func TestIsValidBrackets(t *testing.T) {
	type test struct {
		s    string
		want bool
	}

	tests := map[string]test{
		"t1": {s: "()", want: true},
		"t2": {s: "()[]{}", want: true},
		"t3": {s: "(]", want: false},
		"t4": {s: "([])", want: true},
		"t5": {s: "([)]", want: false},
		"t6": {s: "((", want: false},
	}

	for name, tc := range tests {
		got := IsValidBrackets(tc.s)
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %v, got: %v, want: %v", name, got, tc.want)
			}
		})
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	type test struct {
		strs []string
		want string
	}

	tests := map[string]test{
		"t1": {strs: []string{"flower", "flow", "flight"}, want: "fl"},
		"t2": {strs: []string{"dog", "racecar", "car"}, want: ""},
	}

	for name, tc := range tests {
		got := LongestCommonPrefix(tc.strs)
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %v, got: %v, want: %v", name, got, tc.want)
			}
		})
	}
}

func TestPlusOne(t *testing.T) {
	type test struct {
		digits []int
		want   []int
	}

	tests := map[string]test{
		"t1": {digits: []int{1, 2, 3}, want: []int{1, 2, 4}},
		"t2": {digits: []int{4, 3, 2, 1}, want: []int{4, 3, 2, 2}},
		"t3": {digits: []int{9}, want: []int{1, 0}},
	}

	for name, tc := range tests {
		got := PlusOne(tc.digits)
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %v, got: %v, want: %v", name, got, tc.want)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	type test struct {
		nums  []int
		want1 int
		want2 []int
	}

	tests := map[string]test{
		"t1": {nums: []int{1, 1, 2}, want1: 2, want2: []int{1, 2}},
		"t2": {nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want1: 5, want2: []int{0, 1, 2, 3, 4}},
	}

	for name, tc := range tests {
		got := RemoveDuplicates(tc.nums)
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, tc.want1) {
				t.Errorf("[num error] name: %v, got: %v, want1: %v, want2: %v", name, got, tc.want1, tc.want2)
			} else {
				for i := 0; i < len(tc.want2); i++ {
					if tc.nums[i] != tc.want2[i] {
						t.Errorf("[value error] name: %v, got: %v, want1: %v, want2: %v", name, got, tc.want1, tc.want2)
						break
					}
				}
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type test struct {
		intervals [][]int
		want      [][]int
	}

	tests := map[string]test{
		"t1": {intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		"t2": {intervals: [][]int{{1, 4}, {4, 5}}, want: [][]int{{1, 5}}},
	}

	for name, tc := range tests {
		got := Merge(tc.intervals)
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %v, got: %v, want: %v", name, got, tc.want)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	type test struct {
		nums   []int
		target int
		want   []int
	}

	tests := map[string]test{
		"t1": {nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		"t2": {nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		"t3": {nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for name, tc := range tests {
		got := TwoSum(tc.nums, tc.target)
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %v, got: %v, want: %v", name, got, tc.want)
			}
		})
	}
}
