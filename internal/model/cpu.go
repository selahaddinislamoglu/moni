package model

import "encoding/json"

type CPU struct {
	Usage float64 `json:"usage"`
	Time  int64   `json:"time"`
}

func (c CPU) ToJSON() json.RawMessage {
	data, _ := json.Marshal(c)
	return data
}
