package main

import (
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	file, err := os.Open("test.zip")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// zipファイルをhttpのレスポンスとして返す
	// zipファイルの出力先は単にio.Writerなので、webのレスポンスにも書き込めます
	io.Copy(w, file)
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
