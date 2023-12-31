package proto

import "encoding/json"

type ResponseCode int

type Response struct {
	Code     ResponseCode    `json:"code"`
	Msg      string          `json:"msg"`
	Nickname string          `json:"nickname"`
	Result   string          `json:"result"`
	Data     json.RawMessage `json:"data"`
}
