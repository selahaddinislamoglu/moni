package service

import "github.com/selahaddinislamoglu/moni/backend/model"

type CPU interface {
	GetUsageLastFiveSeconds() (model.CPU, error)
}

type Memory interface {
	GetUsage() (model.Memory, error)
}

type Disk interface {
	GetUsageLastFiveSeconds() (model.Disk, error)
}

type Network interface {
	GetUsageLastFiveSeconds() (model.Network, error)
}
