package main

import (
	"fmt"
)

func main() {
	//areax := area(1, 2)

	y := rect{2, 2}
	z := square{3}
	x := new(square)
	x.width = 4
	a := test()
	fmt.Printf("Area of rect: %v \n", y.area())
	fmt.Printf("Param of rect: %v \n", y.param())
	fmt.Printf("Area of square: %v \n", z.area())
	fmt.Printf("Param of square: %v \n", z.param())
	fmt.Println("################")

	geometry := []geo{y, x, a}
	for _, geox := range geometry {
		fmt.Printf("area is: %v \n", geox.area())
		fmt.Printf("param is: %v \n", geox.param())
	}
	fmt.Println("################")

	fmt.Println(a)

}

func test() *square {

	x := new(square)
	x.width = 10
	return x
}

type geo interface {
	area() int
	param() int
}

type rect struct {
	width  int
	height int
}
type square struct {
	width int
}
type cube struct {
	width int
}

func (x cube) area() int {
	return (x.width * x.width * x.width)
}
func (x *square) area() int {
	return x.width * x.width
}

func (x rect) area() int {
	return x.width * x.height
}

func (x rect) param() int {
	return (x.width + x.height) * 2
}

func (x *square) param() int {
	return (x.width * 4)
}
