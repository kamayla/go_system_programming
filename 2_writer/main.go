package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "http://asci.jp", nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-TEST", "added header")
	req.Write(os.Stdout)
}
