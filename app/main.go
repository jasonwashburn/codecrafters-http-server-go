package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var (
	_ = net.Listen
	_ = os.Exit
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("error accepting connection: ", err.Error())
		os.Exit(1)
	}

	httpVersion := "HTTP/1.1"
	statusCode := "200"
	reasonPhrase := "OK"
	crlf := "\r\n"

	statusLine := fmt.Sprintf("%s %s %s", httpVersion, statusCode, reasonPhrase)

	headers := []string{}

	var sb strings.Builder
	sb.WriteString(statusLine)
	sb.WriteString(crlf)
	for _, header := range headers {
		sb.WriteString(header)
		sb.WriteString(crlf)
	}

	resp := sb.String()

	_, err = conn.Write([]byte(resp))
	if err != nil {
		fmt.Println("failed to write to connection: %w", err)
		os.Exit(1)
	}
}
