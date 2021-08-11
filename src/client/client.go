package client

import (
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
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
	go c.PaintUI()
	<-c.exit
}

func (c *Client) PaintUI() {
	// CursorHide()
	for range time.NewTicker(time.Second).C {
		c.PaintFrame()
	}
}

func (c *Client) PaintFrame() {
	lines := strings.Split(c.Display.Output(), "\n")
	x, y, ok := prepareWindow(lines)
	if !ok {
		return
	}
	for i, l := range lines {
		Gotoxy(x+i, y)
		fmt.Print(l)
	}
}

func prepareWindow(lines []string) (x, y int, ok bool) {
	Clearall()
	Gotoxy(1, 1)
	cw, ch := WinSize()
	mh := len(lines)
	mw := 0
	for _, l := range lines {
		tw := len(l)
		if tw > mw {
			mw = tw
		}
	}
	if mh > ch || mw > cw {
		fmt.Printf(
			"Please change your window size (%d %d) -> (%d+ %d+)",
			cw,
			ch,
			mw,
			mh,
		)
		return 0, 0, false
	}
	return (ch - mh) / 2, (cw - mw) / 2, true
}

func (c *Client) Exit() {
	close(c.exit)
}

func CursorHide() {
	fmt.Println("\033[?25l")
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

func WinSize() (int, int) {
	w, h, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	return w, h
}
