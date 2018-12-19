package main

func main() {
	println("hello")

	defer println("world")

	println("hello1")
	println("hello2")
	println("hello3")

}
