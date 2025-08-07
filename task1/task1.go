package task1

import (
	"slices"
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
