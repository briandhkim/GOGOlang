package main

import "fmt"

func main() {
	gNum := max(3, 1, 4, 5, 2, 14, 21, 3, 4)

	fmt.Println(gNum)
}

func max(nums ...int) int {
	var maxNum int
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}
