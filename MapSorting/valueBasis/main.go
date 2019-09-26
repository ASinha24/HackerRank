package main

import (
	"fmt"
	"sort"
	"strings"
)

type keyVal struct {
	key   string
	value int
}

func main() {
	var laptop map[string]int
	laptop = make(map[string]int)

	laptop["HP"] = 90000
	laptop["Lenovo"] = 25000
	laptop["Acer"] = 50000
	laptop["Apple"] = 150000

	var sliceOfStuct []keyVal

	for k, v := range laptop {
		sliceOfStuct = append(sliceOfStuct, keyVal{k, v})
	}
	fmt.Println("Sorting on the basis of Key")
	print(sortOnkeybasis(sliceOfStuct))

	fmt.Println("Sorting on the basis of value")
	print(sortOnValuebasis(sliceOfStuct))
}

func sortOnkeybasis(slc []keyVal) []keyVal {
	sort.Slice(slc, func(i, j int) bool { return slc[i].key < slc[j].key })
	return slc
}

func sortOnValuebasis(slc []keyVal) []keyVal {
	sort.Slice(slc, func(i, j int) bool { return slc[i].value < slc[j].value })
	return slc
}

func print(slc []keyVal) {
	for _, v := range slc {
		fmt.Printf("%30s%10d\n", v.key, v.value)
	}
	fmt.Println(strings.Repeat("-", 45))
}
