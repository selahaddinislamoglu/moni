package service

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/selahaddinislamoglu/moni/internal/model"
)

type WebsocketService struct {
	Broker Broker
}

func (w *WebsocketService) Setup(broker Broker) {
	w.Broker = broker
}

func (w *WebsocketService) Connect(conn *websocket.Conn) error {
	go func() {
		defer conn.Close()
		subscriber := NewSubscriberService()
		subscriber.Setup(w.Broker)

		var writeLock sync.Mutex
		messageFunc := func(topic string) func(json.RawMessage) {
			return func(message json.RawMessage) {
				data := model.WebsocketResponse{
					Topic: topic,
					Data:  message,
				}
				sent, err := json.Marshal(data)
				if err != nil {
					fmt.Println(err)
					return
				}
				writeLock.Lock()
				conn.WriteMessage(websocket.TextMessage, sent)
				writeLock.Unlock()
			}
		}
		for {
			msgType, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if msgType == websocket.TextMessage {
				var data model.WebsocketRequest

				err := json.Unmarshal(message, &data)
				if err != nil {
					return
				}
				switch data.Action {
				case "subscribe":
					subscriber.Subscribe(data.Topic, messageFunc(data.Topic))
				case "unsubscribe":
					subscriber.Unsubscribe(data.Topic)
				}
			}
		}
	}()
	return nil
}

func NewWebsocketService() Websocket {
	return &WebsocketService{}
}
