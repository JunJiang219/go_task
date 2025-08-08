package task1

import (
	"slices"
	"sort"
	"strconv"

	"github.com/go_task/tools"
)

/*
只出现一次的数字：
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func SingleNumber(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	for num, count := range numMap {
		if count == 1 {
			return num
		}
	}
	return -1
}

// 回文数
func IsPalindrome(x int) bool {
	originStr := strconv.Itoa(x)
	reverseStr := tools.ReverseStrASCII(originStr)

	return (originStr == reverseStr)
}

// 有效的括号
func IsValidBrackets(s string) bool {
	bytes := []byte(s)
	length := len(bytes)
	if length%2 != 0 {
		return false
	}

	stack := tools.Stack{}
	openBrackets := []byte{'(', '[', '{'}
	closeBrackets := []byte{')', ']', '}'}

	for i := 0; i < length; i++ {
		if slices.Contains(openBrackets, bytes[i]) {
			stack.Push(bytes[i])
		} else if slices.Contains(closeBrackets, bytes[i]) {
			val, ok1 := stack.Pop()
			if ok1 {
				v, _ := val.(byte)
				if !tools.IsBracketsPair(v, bytes[i]) {
					return false
				}
			} else {
				return false
			}
		}
	}
	return stack.Length() == 0
}

// 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	runes := [][]rune{}
	for _, v := range strs {
		runes = append(runes, []rune(v))
	}
	count := len(runes)

	minLen := len(runes[0])
	for i := 1; i < len(runes); i++ {
		if minLen > len(runes[i]) {
			minLen = len(runes[i])
		}
	}

	comPre := []rune{}
bigLoop:
	for i := 0; i < minLen; i++ {
		for j := 0; j < count-1; j++ {
			if runes[j][i] != runes[j+1][i] {
				break bigLoop
			}
		}
		comPre = append(comPre, runes[0][i])
	}

	return string(comPre)
}

// 加一
func PlusOne(digits []int) []int {
	length := len(digits)
	tmpSlice := make([]int, length+1, length+1)
	tmpNum, extTag := 0, 0
	for i := length - 1; i >= 0; i-- {
		if i == length-1 {
			tmpNum = digits[i] + 1
		} else {
			tmpNum = digits[i] + extTag
		}

		if tmpNum >= 10 {
			tmpSlice[i+1] = 0
			extTag = 1
		} else {
			tmpSlice[i+1] = tmpNum
			extTag = 0
		}
	}
	tmpSlice[0] = extTag
	if tmpSlice[0] > 0 {
		return tmpSlice
	} else {
		return tmpSlice[1:]
	}
}

// 删除有序数组中的重复项
func RemoveDuplicates(nums []int) int {
	count := 1
	matched := false
Loop1:
	for i := 0; i < len(nums)-1; i++ {
		matched = false
	Loop2:
		for j := i + 1; j < len(nums); j++ {
			if !slices.Contains(nums[:i+1], nums[j]) {
				nums[i+1] = nums[j]
				count++
				matched = true
				break Loop2
			}
		}
		if !matched {
			break Loop1
		}
	}

	return count
}

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// 1. 按区间左端点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 2. 初始化合并结果和当前区间
	res := [][]int{}
	curStart, curEnd := intervals[0][0], intervals[0][1]

	// 3. 遍历并合并重叠区间
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= curEnd {
			// 重叠：更新当前区间右边界（取较大值）
			curEnd = max(curEnd, intervals[i][1])
		} else {
			// 不重叠：保存当前区间并重置
			res = append(res, []int{curStart, curEnd})
			curStart, curEnd = intervals[i][0], intervals[i][1]
		}
	}
	// 4. 添加最后一个合并区间
	res = append(res, []int{curStart, curEnd})
	return res
}
