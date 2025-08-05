package service

import "github.com/selahaddinislamoglu/moni/internal/model"

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

type Authentication interface {
	SetupSecretService(secret Secret)
	Login(request model.LoginRequest) (*model.LoginResponse, error)
}

type Authorization interface {
	SetupSecretService(secret Secret)
	IsAuthorized(token string) bool
}

type Secret interface {
	getJWTsecret() []byte
}
