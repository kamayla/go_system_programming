package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("GET", "http://localhost:8888", nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Accept-Encoding", "gzip")

	// ここのrequestをconn(ソケット)に書き込んだ時点でリクエストは飛ぶ
	request.Write(conn)
	// responseを取得
	response, err := http.ReadResponse(bufio.NewReader(conn), request)

	if err != nil {
		log.Fatal(err)
	}

	dump, err := httputil.DumpResponse(response, true)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(dump))

	defer response.Body.Close()

	if response.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		io.Copy(os.Stdout, reader)
	} else {
		io.Copy(os.Stdout, response.Body)
	}
}
