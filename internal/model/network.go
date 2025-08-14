package model

import "encoding/json"

type Network struct {
	Usage float64 `json:"usage"`
	Time  int64   `json:"time"`
}

func (n Network) ToBytes() []byte {
	data, _ := json.Marshal(n)
	return data
}
