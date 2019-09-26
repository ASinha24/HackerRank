package main

import (
	"fmt"
	"time"
)

//user defined error example
type Myerror struct {
	when time.Time
	what string
}

func (e *Myerror) Error() string {
	return fmt.Sprintf("at %v %s", e.when, e.what)
}

func run() error {
	return &Myerror{time.Now(), "it didn't work"}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
