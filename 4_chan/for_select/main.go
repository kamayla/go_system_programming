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

	/*
		chからデータが転送されてくるまでdefaultの「読み込み中」を表示して
		データが転送してきたらそのデータを表示して待ち受けforループを抜ける
	*/
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
