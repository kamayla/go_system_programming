package main

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	// rand.Readerは無限にランダム文字列を生成するため、io.CopyNで上限バイト数を
	// 設定して1024バイトまでwriterにコピーしたら終了するようにした。
	_, err = io.CopyN(file, rand.Reader, 1024)
	if err != nil {
		log.Fatal(err)
	}
}
