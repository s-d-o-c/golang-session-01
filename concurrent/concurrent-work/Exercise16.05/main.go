package main

import (
	"fmt"
	"log"
)

func greet(ch chan string) {
	fmt.Println("greet():")

	msg := <-ch
	ch <- fmt.Sprintf("Thanks for %s", msg)
	ch <- "Hello Muhammad"
}

func main() {
	ch := make(chan string)

	go greet(ch)

	ch <- "Hello Hari"

	log.Println(<-ch)
	log.Println(<-ch)
}
