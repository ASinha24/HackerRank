package main

import "fmt"

func main() {
	var set = []int32{1, -1, 4}
	fmt.Println(maxCount(set))
}

func maxCount(set []int32) int32 {
	n := uint32(len(set))
	var res, max, temp int32
	res = 0
	max = -999
	//rating = 1
	for i := 0; i < (1 << n); i++ {
		//fmt.Print("{ ")
		res = 0
		for j := 0; j < int(n); j++ {
			if i&(1<<uint32(j)) > 0 {
				res = res + set[j]
				//fmt.Println(" set[j]= ", set[j], " j= ", j, " res = ", res, "rating ", rating)

			}
		}
		//fmt.Println("} ")
		if res > max {
			max = res
			temp = res
		}
	}
	return temp
}
