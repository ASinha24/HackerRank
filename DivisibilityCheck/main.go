package main

import (
	"fmt"
)

func divisibilityCheck(nums []int) {
	max := -999
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			for j := nums[i] * 2; j <= max; j += nums[i] {
				//fmt.Printf(" %d", j)
				for _, v := range nums {
					if v == j {
						count++
						//fmt.Printf("%d *%d *%d\n",v, j, count)
					}
				}
			}
		}
	}
	fmt.Println(max, count)
}
func main() {
	nums := []int{1, 2, 3, 5, 7, 9}
	divisibilityCheck(nums)
}
