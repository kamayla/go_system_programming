package main

import (
	"io"
	"log"
	"os"
)

func main() {
	oldFile, err := os.Open("./old.txt")
	if err != nil {
		log.Fatal(err)
	}

	newFile, err := os.Create("./new.txt")

	if err != nil {
		log.Fatal(err)
	}

	// io.CopyはReaderからWriterにデータを書き込む
	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		log.Fatal("ファイルの書き込みに失敗した")
	}
}
