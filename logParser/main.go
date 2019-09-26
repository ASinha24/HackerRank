package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())
		fmt.Printf("domain: %s - visits: %s\n", fields[0], fields[1])

	}

}
