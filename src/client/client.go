package client

import (
	"fmt"
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
			Clearall()
			Gotoxy(1, 1)
			c.Display.Output()
		}
	}()
}

func (c *Client) Exit() {
	close(c.exit)
}

func Gotoxy(x, y int) {
	fmt.Printf("\033[%d;%df", x, y)
}

func Clearall() {
	fmt.Println("\033[2J")
}

func Clearline() {
	fmt.Println("\033[s\033[K\033[u")
}
