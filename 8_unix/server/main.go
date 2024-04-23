package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path := filepath.Join(os.TempDir(), "unixdomainsocket-sample")
	fmt.Println(path)
	os.Remove(path)
	listen, err := net.Listen("unix", path)

	if err != nil {
		panic(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()

		if err != nil {
			panic(err)
		}

		go func() {
			request, err := http.ReadRequest(bufio.NewReader(conn))

			if err != nil {
				panic(err)
			}

			dump, err := httputil.DumpRequest(request, true)

			if err != nil {
				panic(err)
			}

			fmt.Println(string(dump))

			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       io.NopCloser(strings.NewReader("hello super kamimura")),
			}

			response.Write(conn)
			conn.Close()
		}()
	}
}
