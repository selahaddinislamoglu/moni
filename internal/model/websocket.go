package model

import "encoding/json"

type WebsocketRequest struct {
	Action string `json:"action"`
	Topic  string `json:"topic"`
}

type WebsocketResponse struct {
	Topic string          `json:"topic"`
	Data  json.RawMessage `json:"data"`
}
