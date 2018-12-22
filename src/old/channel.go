package main

import (
	"fmt"
	"strings"
)

func main() {

	phrase := "there are the time that try men's souls. \n"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))

	for _, word := range words {
		fmt.Printf("here is word %v. Adding %v, to channel \n", word, word)
		ch <- word
	}
	close(ch)
	for msg := range ch {
		fmt.Print(msg + " ")
	}

}
