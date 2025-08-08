package main

import (
	"fmt"
	"slices"

	"github.com/go_task/task1"
)

func main() {
	// This is a placeholder for the main function.
	fmt.Printf("【提示】：控制台分别进入各 task 目录，执行命令：go test -v，会输出各测试用例的结果\n\n")

	fmt.Println("begin----------------------- task1 -----------------------------")
	slice1 := []int{2, 2, 1}
	num1 := task1.SingleNumber(slice1)
	fmt.Printf("只出现一次的数字: input: %v, output: %v\n", slice1, num1)

	num2 := 121
	bool2 := task1.IsPalindrome(num2)
	fmt.Printf("回文数: input: %v, output: %v\n", num2, bool2)

	str3 := "()[]{}"
	bool3 := task1.IsValidBrackets(str3)
	fmt.Printf("有效的括号: input: %v, output: %v\n", str3, bool3)

	slice4 := []string{"flower", "flow", "flight"}
	str4 := task1.LongestCommonPrefix(slice4)
	fmt.Printf("最长公共前缀: input: %v, output: %v\n", slice4, str4)

	slice5 := []int{1, 2, 3}
	slice5_1 := task1.PlusOne(slice5)
	fmt.Printf("加一: input: %v, output: %v\n", slice5, slice5_1)

	slice6 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	slice6_1 := slices.Clone(slice6)
	num6 := task1.RemoveDuplicates(slice6)
	fmt.Printf("删除有序数组中的重复项: input: %v, output: %v --- %v\n", slice6_1, num6, slice6)

	slice7 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	slice7_1 := task1.Merge(slice7)
	fmt.Printf("合并区间: input: %v, output: %v\n", slice7, slice7_1)
	fmt.Println("---------------------------- task1 --------------------------end")
}
