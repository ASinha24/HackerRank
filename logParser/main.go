package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var (
		sum     map[string]int
		domains []string
		total   int
		line    int
	)
	sum = make(map[string]int)
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		line++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong Input %v (line  #%d)\n", fields, line)
			return
		}

		domain := fields[0]
		visit, err := strconv.Atoi(fields[1])
		if visit < 0 || err != nil {
			fmt.Printf("wrong Input %v (line  #%d)\n", fields[1], line)
			return
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}

		total += visit
		sum[domain] += visit
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Slice(domains, func(i, j int) bool { return domains[i] < domains[j] })
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
	fmt.Println()
	fmt.Printf("%-30s %10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Println("> err:", err)
	}
}
