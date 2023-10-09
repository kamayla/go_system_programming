package main

import (
	"fmt"
	"time"
)

func main() {
	ch := time.After(10 * time.Second)

	data := <-ch

	fmt.Println(data)
}
