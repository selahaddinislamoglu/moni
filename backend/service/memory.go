package service

import (
	"time"

	"github.com/selahaddinislamoglu/moni/backend/model"
	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryService struct {
}

func NewMemoryService() Memory {
	return &MemoryService{}
}

func (s *MemoryService) GetUsage() (model.Memory, error) {

	var data model.Memory

	v, err := mem.VirtualMemory()
	if err != nil {
		return model.Memory{}, err
	}

	data.Total = v.Total / 1024 / 1024
	data.Used = v.Used / 1024 / 1024
	data.Usage = v.UsedPercent
	data.Time = time.Now().Unix()

	return data, nil
}
