package gomochat

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Client is an interface of Client.
// See NewClient to obtain its implementation.
type Client interface {
	Connect(wsaddr string) error
	Disconnect()
	SendMessage(msg string) error
	AddReceiveMessageListener(lis ReceiveMessageListener)
	RemoveReceiveMessageListener(lis ReceiveMessageListener)
}

// ReceiveMessageListener represents an interface of message listener.
// To receive messages from peer, implement ReceiveMessageListener and
// register the struct using AddReceiveMessageListener method.
type ReceiveMessageListener interface {
	OnReceiveMessage(msg string)
}

type client struct {
	listeners map[ReceiveMessageListener]bool
	wsaddr    string
	conn      *websocket.Conn
	mu        sync.Mutex
}

// NewClient returns and implemetation of Client interface
func NewClient() Client {
	return &client{
		listeners: make(map[ReceiveMessageListener]bool),
	}
}

func (c *client) Connect(wsaddr string) error {
	c.wsaddr = wsaddr
	conn, _, err := websocket.DefaultDialer.Dial(wsaddr, nil)
	if err != nil {
		return fmt.Errorf("failed to establish websocket connection: %s", err.Error())
	}
	c.conn = conn

	go func() {
		defer func() {
			c.conn.Close()
			c.conn = nil
		}()
		for {
			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				log.Printf("receive message: %s", err.Error())
				return
			}
			c.onReceiveMessage(string(msg))
		}
	}()

	return nil
}

func (c *client) Disconnect() {
	if c.conn == nil {
		return
	}
	err := c.conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Printf("disconnect: %s", err.Error())
	}

	_ = c.conn.Close()
	c.conn = nil
}

func (c *client) SendMessage(msg string) error {
	if c.conn == nil {
		if c.wsaddr != "" {
			if err := c.Connect(c.wsaddr); err != nil {
				return fmt.Errorf("connection is not established")
			}
		}
	}
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		return fmt.Errorf("send message: %s", err.Error())
	}
	return nil
}

func (c *client) AddReceiveMessageListener(lis ReceiveMessageListener) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.listeners[lis] = true
}

func (c *client) RemoveReceiveMessageListener(lis ReceiveMessageListener) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.listeners[lis]; ok {
		delete(c.listeners, lis)
	}
}

func (c *client) onReceiveMessage(msg string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k := range c.listeners {
		k.OnReceiveMessage(msg)
	}
}
