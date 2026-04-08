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
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	file, err := io.ReadAll(reader)
	if err != nil {
		return &Request{}, err
	}
	splitFile := strings.Split(string(file), "\r\n")
	if len(splitFile) == 0 || splitFile[0] == "" {
		return &Request{}, fmt.Errorf("empty request line")
	}

	splitLine := strings.Split(splitFile[0], " ")
	if len(splitLine) != 3 {
		return &Request{}, fmt.Errorf("invalid request line: expected 3 parts, got %d", len(splitLine))
	}
	if !strings.HasPrefix(splitLine[2], "HTTP/") {
		return &Request{}, fmt.Errorf("invalid http version: %s", splitLine[2])
	}

	reqLine := RequestLine{}
	reqLine.RequestTarget = splitLine[1]
	reqLine.HttpVersion = strings.TrimPrefix(splitLine[2], "HTTP/")
	reqLine.Method = splitLine[0]
	return &Request{
		RequestLine: reqLine,
	}, nil

}
