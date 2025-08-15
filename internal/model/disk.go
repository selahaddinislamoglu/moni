package model

import "encoding/json"

type Disk struct {
	Usage float64 `json:"usage"`
	Time  int64   `json:"time"`
}

func (d Disk) ToJSON() json.RawMessage {
	data, _ := json.Marshal(d)
	return data
}
