package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortedURL(slc []string) {

	urlm := make(map[string]int)
	count := 0

	for i := 0; i < len(slc); i++ {
		if _, ok := urlm[slc[i]]; !ok {
			count++
		}
		urlm[slc[i]]++
	}

	type kv struct {
		key   string
		value int
	}

	var ss []kv
	for k, v := range urlm {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool { return ss[i].value > ss[j].value })

	for _, v := range ss {
		fmt.Printf("%30s%10d\n", v.key, v.value)
	}
	fmt.Println(strings.Repeat("_", 45))
	fmt.Printf("Total unique URLS %22d\n", count)

}

func main() {
	urls := []string{"google.com", "play.golang.org", "golang.org.com", "outlook.com", "golang.org.com", "golang.org.com", "google.com", "google.com", "google.com"}
	sortedURL(urls)
}
