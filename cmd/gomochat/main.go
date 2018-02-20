package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pankona/gomochat/gomochat"
)

type listener struct{}

func (lis *listener) OnReceiveMessage(msg string) {
	fmt.Printf("message received! %s\n", msg)
}

func main() {
	c := gomochat.NewGomoChatClient()

	lis := &listener{}
	c.AddReceiveMessageListener(lis)
	defer c.RemoveReceiveMessageListener(lis)

	err := c.Connect("127.0.0.1", 8080)
	if err != nil {
		fmt.Printf("failed to establish websocket connection: %s", err.Error())
		os.Exit(1)
	}
	defer c.Disconnect()

	for {
		c.SendMessage("sending test message...")
		<-time.After(2 * time.Second)
	}
}