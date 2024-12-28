package rpc

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"os"
	"strconv"
)

type Server struct {
	ctx context.Context
}

func NewServer(ctx context.Context) *Server {
	return &Server{ctx: ctx}
}

func (s *Server) Listen(handler func(ctx context.Context, req *RequestMessage) (res any, err error)) {
	bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(split)
	for scanner.Scan() {
		msg := scanner.Bytes()
		req, err := DecodeMessage(msg)
		if err != nil {
			log.Printf("Decode error: %s\n\t%s", err, string(msg))
			continue
		}
		res, err := handler(s.ctx, req)
		if err != nil {
			log.Printf("Error at: %s, %s", req.Method, err)
			continue
		}
		if res != nil {
			msg := ResponseMessage{Id: req.Id, Result: res}
			reply := EncodeMessage(msg)
			os.Stdout.Write([]byte(reply))
		}
	}
}

func split(data []byte, atEOF bool) (int, []byte, error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}
	if contentLength > len(content) {
		return 0, nil, nil
	}
	totalLength := len(header) + 4 + contentLength
	return totalLength, content, nil
}
