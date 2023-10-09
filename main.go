package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		defer close(ch)
		fmt.Println("start")
		ch <- "asdf"
		ch <- "asdf"
		ch <- "asdf"
		ch <- "asdf"
		fmt.Println("end")
	}()

	for r := range ch {
		fmt.Println(r)
	}
}
