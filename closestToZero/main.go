package main

import (
	"fmt"
	"math"
)

func closestTozero(nums []int) {
	min := 99999
	for i := 0; i < len(nums); i++ {
		diff := int(math.Abs(float64(nums[i])))
		if diff < min {
			min = diff
		}
	}
	fmt.Println(min)

}
func main() {
	nums := []int{3, -3, 4, -1}
	closestTozero(nums)
}
