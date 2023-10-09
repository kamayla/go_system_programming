package main

import "fmt"

func bufferNotExists() {
	ch := make(chan string)

	go func() {
		defer close(ch)
		/*
			バッファ無しのチャネルの場合は受信側の受信が場合は送信をブロッキングするので
			endが最後に表示される。
		*/
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

func bufferExists() {
	ch := make(chan string, 4)

	go func() {
		defer close(ch)
		/*
			バッファありチャネルの場合は、バッファ数に余裕がある限りはブロックせずに次の処理に
			行く。当処理ではstart,endと処理が走り受信側のループでチャネルの文字列が表示される。
		*/
		fmt.Println("start")
		ch <- "asdf"
		ch <- "asdf"
		ch <- "asdf"
		ch <- "asdf"
		// ここがチャネルで送信した文字列よりも先に表示される
		fmt.Println("end")
	}()

	for r := range ch {
		fmt.Println(r)
	}
}

func main() {
	bufferNotExists()
	fmt.Println("================================================================")
	bufferExists()
}
