package service

type PublisherService struct {
	broker   Broker
	clientID ClientID
}

func NewPublisherService() Publisher {
	return &PublisherService{}
}

func (p *PublisherService) Setup(broker Broker) {
	p.broker = broker
	p.clientID = p.broker.Register()
}

func (p *PublisherService) Publish(topic string, event Event) {
	p.broker.Publish(p.clientID, topic, event)
}
