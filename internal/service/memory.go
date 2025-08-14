package service

import (
	"time"

	"github.com/selahaddinislamoglu/moni/internal/model"
	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryService struct {
	data      model.Memory
	publisher Publisher
}

func NewMemoryService() Memory {
	s := &MemoryService{}
	s.startMonitoring()
	return s
}

func (s *MemoryService) startMonitoring() {
	go func() {
		time.Sleep(5 * time.Second)
		for {
			start := time.Now()
			data, err := s.GetUsage()
			if err != nil {
				continue
			}
			s.data = data
			s.publisher.Publish(MemoryTopic, s.data)
			time.Sleep(5*time.Second - time.Since(start))
		}
	}()
}

func (s *MemoryService) Setup(Publisher Publisher) {
	s.publisher = Publisher
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
