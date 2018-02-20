package gomochat

import (
	"fmt"
	"sync"
)

type gomoChatClient struct {
	listeners map[ReceiveMessageListener]bool
	mu        sync.Mutex
}

//export
type GomoChatClient interface {
	Connect(ipaddress string, port int) error
	Disconnect()
	SendMessage(msg string)
	AddReceiveMessageListener(lis ReceiveMessageListener)
	RemoveReceiveMessageListener(lis ReceiveMessageListener)
}

//export
type ReceiveMessageListener interface {
	OnReceiveMessage(msg string)
}

//export
func NewGomoChatClient() GomoChatClient {
	return &gomoChatClient{
		listeners: make(map[ReceiveMessageListener]bool),
	}
}

//export
func (c *gomoChatClient) Connect(ipaddress string, port int) error {
	fmt.Printf("TODO IMPLEMENT\n")
	return nil
}

//export
func (c *gomoChatClient) Disconnect() {
	fmt.Printf("TODO IMPLEMENT\n")
}

//export
func (c *gomoChatClient) SendMessage(msg string) {
	fmt.Printf("TODO IMPLEMENT\n")
}

//export
func (c *gomoChatClient) AddReceiveMessageListener(lis ReceiveMessageListener) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.listeners[lis] = true
}

//export
func (c *gomoChatClient) RemoveReceiveMessageListener(lis ReceiveMessageListener) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.listeners[lis]; ok {
		delete(c.listeners, lis)
	}
}

func (c *gomoChatClient) onReceiveMessage(msg string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k := range c.listeners {
		k.OnReceiveMessage(msg)
	}
}
