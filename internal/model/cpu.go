package model

import "encoding/json"

type CPU struct {
	Usage float64 `json:"usage"`
	Time  int64   `json:"time"`
}

func (c CPU) ToBytes() []byte {
	data, _ := json.Marshal(c)
	return data
}
