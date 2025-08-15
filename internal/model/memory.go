package model

import "encoding/json"

type Memory struct {
	Total uint64  `json:"total"`
	Used  uint64  `json:"used"`
	Usage float64 `json:"usage"`
	Time  int64   `json:"time"`
}

func (m Memory) ToJSON() json.RawMessage {
	data, _ := json.Marshal(m)
	return data
}
