package main

import (
	"fmt"
)

func main() {
	bnt := MakeButton()
	handler1 := make(chan string)
	handler2 := make(chan string)

	bnt.AddEventListener("click", handler1)
	bnt.AddEventListener("click", handler2)

	go func() {
		for {
			msg := <-handler1
			fmt.Println("handler 1 message:", msg)
		}
	}()
	go func() {
		for {
			msg := <-handler2
			fmt.Println("handler 2 message:", msg)
		}
	}()

	bnt.TriggerEvent("click", "button is clicked!")
	bnt.RemoveEventListener("click", handler2)
	bnt.TriggerEvent("click", "Buttn is clicked again!!")
	fmt.Scanln()

}

type Button struct {
	eventListeners map[string][]chan string
}

func MakeButton() *Button {
	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
}

func (this *Button) AddEventListener(event string, responseChannel chan string) {
	if _, present := this.eventListeners[event]; present {
		this.eventListeners[event] = append(this.eventListeners[event], responseChannel)
	} else {
		this.eventListeners[event] = []chan string{responseChannel}
	}
}

func (this *Button) RemoveEventListener(event string, listenerChannel chan string) {
	if _, present := this.eventListeners[event]; present {
		for idx, _ := range this.eventListeners[event] {
			if this.eventListeners[event][idx] == listenerChannel {
				this.eventListeners[event] = append(this.eventListeners[event][:idx], this.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

func (this *Button) TriggerEvent(event, response string) {
	if _, present := this.eventListeners[event]; present {
		for _, handler := range this.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}
