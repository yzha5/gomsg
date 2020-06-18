package gomsg

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestMessage_Code(t *testing.T) {
	code := NewMsg(123, "test message").Code()
	fmt.Println(reflect.TypeOf(code)) //int
	fmt.Println(code)                 //123
}

func TestMessage_Msg(t *testing.T) {
	msg := NewMsg("123", "test message").Msg()
	fmt.Println(reflect.TypeOf(msg)) //string
	fmt.Println(msg)                 //test message
}

func TestMessage_Err(t *testing.T) {
	err := errors.New("test error message")
	msg := NewMsgByError(321, err)
	fmt.Println(msg.Code()) //321
	fmt.Println(msg.Msg())  //test error message
	fmt.Println(msg.Err())  //test error message
}

func TestMessage_Print(t *testing.T) {
	msg := NewMsg(434, "abc def ghi jkl")
	msg.Print()  //status:434 text:abc def ghi jkl
	msg.LPrint() //2020/06/17 20:01:21 status:434 text:abc def ghi jkl
}

func TestUnMarshal(t *testing.T) {
	message0 := UnMarshal("{\"Status\":321,\"Text\":\"test error message\"}")
	fmt.Println(message0.Status)
	fmt.Println(message0.Text)

	message1 := UnMarshal("123")
	fmt.Println(message1.Status)
	fmt.Println(message1.Text)
}

func TestMarshal(t *testing.T) {
	err := Marshal("TestStatus", "abc")
	fmt.Println(err)
}
