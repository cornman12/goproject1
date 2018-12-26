package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	po := new(PurchaseOrder)
	po.Value = 32.1

	SavePO(po, false).Then(func(obj interface{}) error {
		po := obj.(*PurchaseOrder)
		fmt.Printf("purchase order saved with id %d \n", po.Number)
		return nil
	}, func(err error) {
		fmt.Printf("Failed to save Purchase order: " + err.Error() + "\n")
	}).Then(func(obj interface{}) error {
		fmt.Println("Second Promise Success")
		return nil
	}, func(err error) {
		fmt.Println("Second Promise Failed: " + err.Error())
	})
	fmt.Scanln()

}

type PurchaseOrder struct {
	Number int
	Value  float64
}

func SavePO(po *PurchaseOrder, shouldFail bool) *Promise {
	result := new(Promise)
	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		time.Sleep(2 * time.Second)
		if shouldFail {
			result.failureChannel <- errors.New("Failed to save purchase Order")
		} else {
			po.Number = 1234
			result.successChannel <- po
		}
	}()
	return result

}

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

func (this *Promise) Then(sucess func(interface{}) error, failure func(error)) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)
	timeout := time.After(1 * time.Second)

	go func() {
		select {
		case obj := <-this.successChannel:
			newErr := sucess(obj)
			if newErr == nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newErr
			}
		case err := <-this.failureChannel:
			failure(err)
			result.failureChannel <- err
		case <-timeout:
			failure(errors.New("shit just got timeout"))
		}
	}()
	return result

}
