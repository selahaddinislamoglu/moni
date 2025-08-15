package service

import "encoding/json"

type SubscriberService struct {
	broker   Broker
	clientID ClientID
}

func NewSubscriberService() Subscriber {
	return &SubscriberService{}
}

func (s *SubscriberService) Setup(broker Broker) {
	s.broker = broker
	s.clientID = s.broker.Register()
}

func (s *SubscriberService) Subscribe(topic string, handler func(message json.RawMessage)) {
	s.broker.Subscribe(s.clientID, topic, handler)
}

func (s *SubscriberService) Unsubscribe(topic string) {
	s.broker.Unsubscribe(s.clientID, topic)
}
