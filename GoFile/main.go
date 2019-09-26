package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hello..... This is Alka")

	defer file.Close()

	stream, err := ioutil.ReadFile("file.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(stream))
}
