package request

import (
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}
type Request struct {
	RequestLine RequestLine
	ParserState int
}

// func (r *Request) parse(data []byte) (int, error)

func RequestFromReader(reader io.Reader) (*Request, int, error) {
	file, err := io.ReadAll(reader)
	if err != nil {
		return &Request{}, 0, err
	}
	if !strings.Contains(string(file), "\r\n") {
		return &Request{}, 0, fmt.Errorf("no \r\n yet")
	}
	splitFile := strings.Split(string(file), "\r\n")
	if len(splitFile) == 0 || splitFile[0] == "" {
		return &Request{}, 0, fmt.Errorf("empty request line")
	}
	splitLine := strings.Split(splitFile[0], " ")
	reqLine, err := parseRequestLine(splitLine)
	if err != nil {
		return &Request{}, 0, err
	}
	consumed := len(splitFile[0]) + len("\r\n")

	return &Request{
		RequestLine: *reqLine,
	}, consumed, nil

}

func parseRequestLine(splitLine []string) (*RequestLine, error) {
	if len(splitLine) != 3 {
		return &RequestLine{}, fmt.Errorf("invalid request line: expected 3 parts, got %d", len(splitLine))
	}
	if !strings.HasPrefix(splitLine[2], "HTTP/") {
		return &RequestLine{}, fmt.Errorf("invalid http version: %s", splitLine[2])
	}

	reqLine := RequestLine{}
	reqLine.RequestTarget = splitLine[1]
	reqLine.HttpVersion = strings.TrimPrefix(splitLine[2], "HTTP/")
	reqLine.Method = splitLine[0]

	return &reqLine, nil

}
