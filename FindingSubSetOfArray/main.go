package main

import "fmt"

func main() {
	var set = []int32{1, 2, 3}
	subset(set)
}

func subset(set []int32) {

	n := uint32(len(set))
	for i := 0; i < (1 << n); i++ {
		fmt.Printf("{ ")
		for j := 0; j < int(n); j++ {
			if i&(1<<uint32(j)) > 0 {
				fmt.Print(set[j], " ")
			}
		}
		fmt.Println("}")
	}

}
