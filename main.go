package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Print(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	numsNew := make(map[int]int)
	for idx, elem := range nums {
		need := target - elem
		if value, inMap := numsNew[need]; inMap {
			return []int{value, idx}
		}
		numsNew[elem] = idx
	}
	return []int{}
}
