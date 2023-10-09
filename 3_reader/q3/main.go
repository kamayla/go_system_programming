package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("test.zip")

	if err != nil {
		log.Fatal(err)
	}
	zipWriter := zip.NewWriter(file)

	defer zipWriter.Close()

	aWriter, err := zipWriter.Create("a.txt")

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(aWriter, strings.NewReader("hogehoge"))

	bWriter, err := zipWriter.Create("b.txt")

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(bWriter, strings.NewReader("fugafuga"))

}
