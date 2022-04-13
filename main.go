package main

import (
	"LeetCode/problems"
	"fmt"
)

func main() {
	//str := "12一"
	//fmt.Println(utf8.RuneCountInString(str)) // 3
	//fmt.Println(len(str))                    // 5
	//s := " 1 2  3"
	//fmt.Println(strings.Fields(s))

	res := problems.TwoSumTarget([]int{1, 1, 1, 2, 2, 3, 3}, 4)
	fmt.Println(res)
}
