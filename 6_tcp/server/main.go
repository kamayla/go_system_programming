package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		// AcceptでClientからのConnection確立を待ち受ける
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			// リクエストを読む
			request, err := http.ReadRequest(bufio.NewReader(conn))

			if err != nil {
				log.Fatal(err)
			}

			dump, err := httputil.DumpRequest(request, true)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(dump))

			// Responseを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 1,
				Body:       io.NopCloser(strings.NewReader("<h1>Hello World</h1>")),
			}

			// conn(ソケット)にresponseの内容を書き込んでいる
			response.Write(conn)

			conn.Close()

		}()
	}
}
