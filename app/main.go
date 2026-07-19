package main

import (
	"bufio"
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

const (
	CRLF = "\r\n"
)

type httpRequest struct {
	method      string
	target      string
	httpVersion string
}

type httpResponse struct {
	httpVersion  string
	statusCode   string
	reasonPhrase string
}

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

	// request phase
	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("requestLine does not end in crlf: ", requestLine)
		os.Exit(1)
	}
	requestLine = strings.TrimSpace(requestLine)
	requestParts := strings.Split(requestLine, " ")
	if len(requestParts) != 3 {
		fmt.Println("request line does not contain exactly 3 parts: ", requestLine)
	}

	req := httpRequest{
		method:      requestParts[0],
		target:      requestParts[1],
		httpVersion: requestParts[2],
	}

	// response phase
	resp := httpResponse{}

	if req.target == "/" {
		resp.httpVersion = "HTTP/1.1"
		resp.statusCode = "200"
		resp.reasonPhrase = "OK"
	} else {
		resp.httpVersion = "HTTP/1.1"
		resp.statusCode = "404"
		resp.reasonPhrase = "Not Found"
	}

	statusLine := fmt.Sprintf("%s %s %s", resp.httpVersion, resp.statusCode, resp.reasonPhrase)

	headers := []string{}
	var sb strings.Builder
	sb.WriteString(statusLine)
	sb.WriteString(CRLF)
	for _, header := range headers {
		sb.WriteString(header)
		sb.WriteString(CRLF)
	}
	sb.WriteString(CRLF)

	craftedResponse := sb.String()

	_, err = conn.Write([]byte(craftedResponse))
	if err != nil {
		fmt.Println("failed to write to connection: %w", err)
		os.Exit(1)
	}
}
