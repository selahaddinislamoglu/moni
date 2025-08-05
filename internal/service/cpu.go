package service

import (
	"fmt"
	"time"

	"github.com/selahaddinislamoglu/moni/internal/model"
	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUService struct {
	data model.CPU
}

func NewCPUService() CPU {

	cpu := &CPUService{}
	cpu.startMonitoring()
	return cpu
}

func (c *CPUService) startMonitoring() {
	go func() {
		for {
			percent, err := cpu.Percent(time.Second*5, false)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			c.data.Usage = percent[0]
			c.data.Time = time.Now().Unix()
		}
	}()
}

func (c *CPUService) GetUsageLastFiveSeconds() (model.CPU, error) {

	return c.data, nil
}
