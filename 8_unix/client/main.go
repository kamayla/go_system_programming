package main

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	conn, err := net.Dial("unix",
		filepath.Join(os.TempDir(), "unixdomainsocket-sample"),
	)

	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("get", "http://localhost:8888", nil)

	if err != nil {
		panic(err)
	}

	request.Write(conn)

	response, err := http.ReadResponse(bufio.NewReader(conn), request)

	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)

	io.Copy(os.Stdout, strings.NewReader(string(dump)))
}
