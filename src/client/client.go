package client

import (
	"time"
)

type Client struct {
	exit     chan bool
	Display  *Display
	Operator *Operator
}

func New() *Client {
	return &Client{
		exit:     make(chan bool),
		Display:  &Display{},
		Operator: &Operator{},
	}
}

func (c *Client) Start() {
	go c.Operator.Scan()
	<-c.exit
}

func (c *Client) Refresh() {
	go func() {
		for range time.NewTicker(time.Second).C {
			c.Display.Output()
			c.Operator.Output()
		}
	}()
}

func (c *Client) Exit() {
	close(c.exit)
}
