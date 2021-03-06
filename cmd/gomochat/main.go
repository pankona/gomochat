package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pankona/gomochat"
)

type listener struct{}

func (lis *listener) OnReceiveMessage(msg string) {
	fmt.Printf("message received! %s\n", msg)
}

func main() {
	c := gomochat.NewClient()

	lis := &listener{}
	c.AddReceiveMessageListener(lis)
	defer c.RemoveReceiveMessageListener(lis)

	err := c.Connect("ws://127.0.0.1:8080/ws")
	if err != nil {
		fmt.Printf("failed to establish websocket connection: %s\n", err.Error())
		os.Exit(1)
	}
	defer c.Disconnect()

	for {
		err = c.SendMessage("sending test message...")
		if err != nil {
			fmt.Printf("sendMessage returned error: %s\n", err.Error())
		}

		<-time.After(1 * time.Second)
	}
}
