package main

import (
	"fmt"
)

func main() {
	po := new(PurchaseOrder)
	po.Value = 32.12
	ch := make(chan *PurchaseOrder)

	go SavePO(po, ch)
	newPO := <-ch
	fmt.Printf("PO: %v \n", newPO)

}

type PurchaseOrder struct {
	Number int
	Value  float64
}

func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.Number = 1234
	callback <- po
}
