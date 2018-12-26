package main

import (
	"fmt"
)

func main() {
	msgCh := make(chan Message, 2)
	errCh := make(chan FailedMessage, 2)

	msg := Message{
		To: []string{"tomlee@gmail.com",
			"tomx@gmail.com"},
		From:    "toomlee@gmail.com",
		Content: "Here i am",
	}
	msgCh <- msg

	select {
	case recievedMsg := <-msgCh:
		fmt.Println(recievedMsg)
	case recievedError := <-errCh:
		fmt.Println(recievedError)
	default:
		fmt.Println("No more message")

	}

	/*
		failedMessage := FailedMessage{
			ErrorMessage:    "Message is errored out",
			OriginalMessage: Message{},
		}

		jsonobject, _ := json.Marshal(msg)

		fmt.Println(string(jsonobject))

		msgCh <- msg
		errCh <- failedMessage
		fmt.Println(<-msgCh)
		fmt.Println(<-errCh)
	*/

}

type Message struct {
	To      []string
	From    string
	Content string
}
type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}
