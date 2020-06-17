package gomsg

import (
	"errors"
	"fmt"
	"log"
)

type C interface {
	Code() interface{}
	Msg() string
	Err() error
	Print()
	LPrint()
}

type Message struct {
	Status interface{}
	Text   string
}

func (c *Message) Code() interface{} {
	return c.Status
}

func (c *Message) Msg() string {
	return c.Text
}

func (c *Message) Err() error {
	return errors.New(c.Text)
}

func (c *Message) Print() {
	fmt.Printf("%sstatus%s%s%v%s %stext%s%s%s%s\n",
		blueBg, reset, cyanBg, c.Status, reset, blueBg, reset, cyanBg, c.Text, reset)
}

func (c *Message) LPrint() {
	log.Printf("%sstatus%s%s%v%s %stext%s%s%s%s\n",
		blueBg, reset, cyanBg, c.Status, reset, blueBg, reset, cyanBg, c.Text, reset)
}

func New(status interface{}, msg string) C {
	return &Message{
		Status: status,
		Text:   msg,
	}
}

func NewWithError(status interface{}, err error) C {
	return &Message{
		Status: status,
		Text:   err.Error(),
	}
}
