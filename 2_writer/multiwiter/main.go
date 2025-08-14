package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Encoding", "gzip")
	source := map[string]string{
		"Hello": "worlda",
	}

	jsonBytes, err := json.Marshal(source)
	if err != nil {
		log.Fatalln(err)
	}

	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()
	multiWriter := io.MultiWriter(gzipWriter, os.Stdout)
	io.WriteString(multiWriter, string(jsonBytes))
}
func main() {
	http.HandleFunc("/", handleFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
