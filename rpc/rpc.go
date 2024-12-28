package rpc

import (
	"encoding/json"
	"fmt"
)

type RequestMessage struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Id     int             `json:"id"`
}

type ResponseMessage struct {
	Id     int            `json:"id"`
	Result *any           `json:"result,omitempty"`
	Error  *ResponseError `json:"error,omitempty"`
}

type ResponseError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Data    *any      `json:"data,omitempty"`
}

type ErrorCode int

const (
	RequestFailedCode ErrorCode = -32803
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMessage(msg []byte) (*RequestMessage, error) {
	var message RequestMessage
	if err := json.Unmarshal(msg, &message); err != nil {
		return nil, err
	}
	return &message, nil
}
