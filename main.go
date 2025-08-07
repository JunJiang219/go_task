package main

import (
	"fmt"

	"github.com/go_task/task1"
)

func main() {
	// This is a placeholder for the main function.
	fmt.Printf("【提示】：控制台分别进入各 task 目录，执行命令：go test -v，会输出各测试用例的结果\n\n")

	fmt.Println("begin----------------------- task1 -----------------------------")
	slice1 := []int{2, 2, 1}
	num1 := task1.SingleNumber(slice1)
	fmt.Printf("只出现一次的数字: input: %v, output: %v\n", slice1, num1)

	num1 = 121
	bool1 := task1.IsPalindrome(num1)
	fmt.Printf("回文数: input: %v, output: %v\n", num1, bool1)

	str1 := "()[]{}"
	bool1 = task1.IsValidBrackets(str1)
	fmt.Printf("有效的括号: input: %v, output: %v\n", str1, bool1)

	slice2 := []string{"flower", "flow", "flight"}
	str1 = task1.LongestCommonPrefix(slice2)
	fmt.Printf("最长公共前缀: input: %v, output: %v\n", slice2, str1)
	fmt.Println("---------------------------- task1 --------------------------end")
}
