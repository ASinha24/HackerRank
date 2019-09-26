package main

import "fmt"

func main() {
	var arr = []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	count := [6]int32{0, 0, 0, 0, 0, 0}
	var maxval, maxpos int32
	maxval, maxpos = 0, 0

	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	for i, v := range count {
		fmt.Println("index ", i, "value ", v)
	}

	for i := 0; i < len(count); i++ {
		if count[i] > maxval {
			maxval = count[i]
			maxpos = int32(i)
			fmt.Println("maxval ", maxval, "pos ", maxpos)
		}
	}

}
