package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Create("chin.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buffer := bufio.NewWriter(file)
	buffer.WriteString("hogehoge")
	time.Sleep(2 * time.Second)
	buffer.Flush()

	buffer.WriteString("hogehoge")
	buffer.WriteString("hogehoge")
	buffer.WriteString("hogehoge")

	time.Sleep(2 * time.Second)
	buffer.Flush()
}
