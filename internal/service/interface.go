package service

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/selahaddinislamoglu/moni/internal/model"
)

const (
	CPUTopic     = "cpu"
	MemoryTopic  = "memory"
	DiskTopic    = "disk"
	NetworkTopic = "network"
)

type CPU interface {
	Setup(publisher Publisher)
	GetUsageLastFiveSeconds() (model.CPU, error)
}

type Memory interface {
	Setup(publisher Publisher)
	GetUsage() (model.Memory, error)
}

type Disk interface {
	Setup(publisher Publisher)
	GetUsageLastFiveSeconds() (model.Disk, error)
}

type Network interface {
	Setup(publisher Publisher)
	GetUsageLastFiveSeconds() (model.Network, error)
}

type Authentication interface {
	Setup(secret Secret)
	Login(request model.LoginRequest) (*model.LoginResponse, error)
}

type Authorization interface {
	Setup(secret Secret)
	IsAuthorized(token string) bool
}

type Secret interface {
	getJWTsecret() []byte
}

type ClientID string

type Event interface {
	ToJSON() json.RawMessage
}

type Broker interface {
	Register() ClientID
	Unregister(id ClientID)
	Subscribe(id ClientID, topic string, handler func(message json.RawMessage))
	Unsubscribe(id ClientID, topic string)
	Publish(id ClientID, topic string, event Event)
}

type Subscriber interface {
	Setup(broker Broker)
	Subscribe(topic string, handler func(message json.RawMessage))
	Unsubscribe(topic string)
}

type Publisher interface {
	Setup(broker Broker)
	Publish(topic string, event Event)
}

type Websocket interface {
	Setup(broker Broker)
	Connect(conn *websocket.Conn) error
}
