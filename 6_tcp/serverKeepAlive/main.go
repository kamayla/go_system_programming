package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func isGzipAcceptable(request *http.Request) bool {
	return strings.Index(strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

func processSession(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept %v\n", conn.RemoteAddr())

	// Accept後のconnで何度も応答を返すためのるーぷ
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		// リクエストを読む
		request, err := http.ReadRequest(bufio.NewReader(conn))

		// timeoutもしくはソケットのクローズ時は終了
		// それ以外はエラーになる
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Time Out")
				break
			} else if err == io.EOF {
				break
			}
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
			Header:     make(http.Header),
		}

		if isGzipAcceptable(request) {
			content := "<h1>Hello World (gzip)</h1>"
			var buf bytes.Buffer
			writer := gzip.NewWriter(&buf)
			io.WriteString(writer, content)
			writer.Close()
			response.Body = io.NopCloser(&buf)
			response.ContentLength = int64(buf.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "<h1>Hello World</h1>"
			response.ContentLength = int64(len(content))
			response.Body = io.NopCloser(strings.NewReader(content))
		}

		// conn(ソケット)にresponseの内容を書き込んでいる
		response.Write(conn)
	}
}
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

		go processSession(conn)
	}
}
