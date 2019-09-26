package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	argument := os.Args
	if len(argument) < 2 {
		fmt.Printf("Usages main.go colnumber file1 [file2.....filen]")
	}

	temp, err := strconv.Atoi(argument[1])

	if err != nil {
		fmt.Println("Invalid column number")
		os.Exit(1)
	}
	column := temp
	if column < 0 {
		fmt.Println("Please enter valid column number")
	}

	for _, filename := range argument[2:] {
		fmt.Println("filename", filename)

		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file ", err)
			continue
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("error reading file ", err)
			}

			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}

		}
	}
}
