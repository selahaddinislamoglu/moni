package service

import (
	"fmt"
	"time"

	"github.com/selahaddinislamoglu/moni/internal/model"
	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUService struct {
	data      model.CPU
	publisher Publisher
}

func NewCPUService() CPU {

	cpu := &CPUService{}
	cpu.startMonitoring()
	return cpu
}

func (c *CPUService) Setup(Publisher Publisher) {
	c.publisher = Publisher
}

func (c *CPUService) startMonitoring() {
	go func() {
		var readTime int64 = time.Now().UnixMilli()
		for {
			percent, err := cpu.Percent(time.Second*5-time.Since(time.UnixMilli(readTime)), false)
			readTime = time.Now().UnixMilli()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			c.data.Usage = percent[0]
			c.data.Time = time.Now().Unix()
			c.publisher.Publish(CPUTopic, c.data)
		}
	}()
}

func (c *CPUService) GetUsageLastFiveSeconds() (model.CPU, error) {

	return c.data, nil
}
