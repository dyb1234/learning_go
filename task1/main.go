package main

import (
	"fmt"
	"task1/questions"
)

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(questions.SingleNumber(nums))

	x := 121
	fmt.Println(questions.IsPalindrome(x))
	x = 123
	fmt.Println(questions.IsPalindrome(x))

	s := "[[]]"
	fmt.Println(questions.IsValid(s))
	s = "(("
	fmt.Println(questions.IsValid(s))

	words := []string{"flower", "flow", "flight"}
	fmt.Println(questions.LongestCommonPrefix(words))

	nums = []int{1, 2, 3}
	fmt.Println(questions.PlusOne(nums))

	nums = []int{1, 1, 2}
	fmt.Println(nums)
	fmt.Println(questions.RemoveDuplicates(nums))

	intervals := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	result := questions.Merge(intervals)
	fmt.Println(result)

	nums = []int{2, 7, 11, 15}
	target := 9
	fmt.Println(questions.TwoSum(nums, target))
}
