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
