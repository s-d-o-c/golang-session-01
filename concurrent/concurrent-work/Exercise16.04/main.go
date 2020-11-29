package main

import (
	"fmt"
	"log"
)

func greet(ch chan string) {
	fmt.Println("greet(): ch <- \"Hello\"")
	ch <- "Hello"
}

func main() {
	ch := make(chan string)

	go greet(ch)

	log.Println(<-ch)
}
