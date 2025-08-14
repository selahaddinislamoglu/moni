package service

import (
	"fmt"
	"time"
)

type brokerClient struct {
	ID     ClientID
	Topics map[string]func(message []byte)
}

type BrokerService struct {
	clients       map[ClientID]*brokerClient
	topicChannels map[string]map[ClientID]chan []byte
}

func NewBrokerService() Broker {
	return &BrokerService{
		clients:       make(map[ClientID]*brokerClient),
		topicChannels: make(map[string]map[ClientID]chan []byte),
	}
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

func (b *BrokerService) Register() ClientID {
	id := fmt.Sprintf("%d-%s", time.Now().UnixNano(), randomString(6))
	client := &brokerClient{
		ID:     ClientID(id),
		Topics: make(map[string]func(message []byte)),
	}
	b.clients[client.ID] = client
	return client.ID
}

func (b *BrokerService) Unregister(id ClientID) {
	client, exists := b.clients[id]
	if !exists {
		return
	}
	for topic := range client.Topics {
		b.Unsubscribe(id, topic)
	}
	delete(b.clients, id)
}

func (b *BrokerService) Subscribe(id ClientID, topic string, handler func(message []byte)) {
	client, exists := b.clients[id]
	if !exists {
		return
	}
	client.Topics[topic] = handler

	if _, exists := b.topicChannels[topic]; !exists {
		b.topicChannels[topic] = make(map[ClientID]chan []byte)
	}

	ch := make(chan []byte)
	b.topicChannels[topic][id] = ch

	go func() {
		for {
			message, ok := <-ch
			if !ok {
				return
			}
			handler(message)
		}
	}()
}

func (b *BrokerService) Unsubscribe(id ClientID, topic string) {
	client, exists := b.clients[id]
	if !exists {
		return
	}
	ch, exists := b.topicChannels[topic][id]
	if exists {
		close(ch)
	}
	delete(b.topicChannels[topic], id)
	delete(client.Topics, topic)
}

func (b *BrokerService) Publish(id ClientID, topic string, event Event) {
	_, exists := b.clients[id]
	if !exists {
		return
	}
	channels, exists := b.topicChannels[topic]
	if !exists {
		return
	}
	go func() {
		message := event.ToBytes()
		if len(message) == 0 {
			fmt.Printf("No message to publish for topic %s, client %s\n", topic, id)
			return
		}
		for _, ch := range channels {
			select {
			case ch <- message:
			default:
				fmt.Printf("Message dropped for topic %s, client %s\n", topic, id)
			}
		}
	}()
}
