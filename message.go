package gomsg

import (
	"encoding/json"
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

func (m *Message) Code() interface{} {
	return m.Status
}

func (m *Message) Msg() string {
	return m.Text
}

func (m *Message) Err() error {
	mJson, _ := json.Marshal(m)
	return errors.New(string(mJson))
}

func (m *Message) Print() {
	fmt.Printf("%sstatus%s%s%v%s %stext%s%s%s%s\n",
		blueBg, reset, cyanBg, m.Status, reset, blueBg, reset, cyanBg, m.Text, reset)
}

func (m *Message) LPrint() {
	log.Printf("%sstatus%s%s%v%s %stext%s%s%s%s\n",
		blueBg, reset, cyanBg, m.Status, reset, blueBg, reset, cyanBg, m.Text, reset)
}

func Marshal(status interface{}, msg string) error {
	return NewMsg(status, msg).Err()
}

func UnMarshal(msg string) *Message {
	m := new(Message)
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return &Message{
			Status: "Unknown",
			Text:   msg,
		}
	}
	return m
}

func NewMsg(status interface{}, msg string) *Message {
	return &Message{
		Status: status,
		Text:   msg,
	}
}

func NewMsgByError(status interface{}, err error) *Message {
	return &Message{
		Status: status,
		Text:   err.Error(),
	}
}
