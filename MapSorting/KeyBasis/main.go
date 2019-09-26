package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var url [10]string
	//url := []string{"golang.org.com", "wikipedia.com", "golang.org.com", "wikipedia.com", "golang.org.com", "wikipedia.com", "annotation.com"}
	fmt.Println("enter 5 URLs")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%s\n", &url[i])
	}
	fmt.Println(strings.Repeat("*", 45))
	urlsort(url[:5])
	fmt.Println(strings.Repeat("_", 45))
}

func urlsort(slc []string) {
	var domains []string
	mapURL := make(map[string]int)
	for i := 0; i < len(slc); i++ {
		if _, ok := mapURL[slc[i]]; !ok {
			domains = append(domains, slc[i])
		}
		mapURL[slc[i]]++
	}
	sort.Strings(domains)
	fmt.Println(domains)
	for _, domain := range domains {
		fmt.Printf("%30s%10d\n", domain, mapURL[domain])
	}

}
