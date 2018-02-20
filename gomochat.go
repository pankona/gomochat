package gomochat

import (
	"fmt"
	"sync"
)

type client struct {
	listeners map[ReceiveMessageListener]bool
	mu        sync.Mutex
}

// Client is an interface of Client.
// See NewClient to obtain its implementation.
type Client interface {
	Connect(ipaddress string, port int) error
	Disconnect()
	SendMessage(msg string)
	AddReceiveMessageListener(lis ReceiveMessageListener)
	RemoveReceiveMessageListener(lis ReceiveMessageListener)
}

// ReceiveMessageListener represents an interface of message listener.
// To receive messages from peer, implement ReceiveMessageListener and
// register the struct using AddReceiveMessageListener method.
type ReceiveMessageListener interface {
	OnReceiveMessage(msg string)
}

// NewClient returns and implemetation of Client interface
func NewClient() Client {
	return &client{
		listeners: make(map[ReceiveMessageListener]bool),
	}
}

func (c *client) Connect(ipaddress string, port int) error {
	fmt.Printf("(Connect) TODO IMPLEMENT\n")
	return nil
}

func (c *client) Disconnect() {
	fmt.Printf("(Disconnect) TODO IMPLEMENT\n")
}

func (c *client) SendMessage(msg string) {
	fmt.Printf("(SendMessage) TODO IMPLEMENT\n")
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

// just comment out this function since this will be used later.
// comment out to hide from gometalinter.
/*
func (c *client) onReceiveMessage(msg string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k := range c.listeners {
		k.OnReceiveMessage(msg)
	}
}
*/
