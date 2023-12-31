package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		ch <- "転送データ"
	}()

LOOP:
	for {
		time.Sleep(1 * time.Second)
		select {
		case data := <-ch:
			fmt.Println(data)
			break LOOP
		default:
			fmt.Println("読み込み中")
		}
	}
}
