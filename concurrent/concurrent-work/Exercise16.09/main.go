package main

import (
	"fmt"
	"log"
)

func readThem(in, out chan string) {
	fmt.Println("readThem(): in=", in, " out=", out)
	for i := range in {
		log.Println(i)
	}
	out <- "done"
}

func main() {
	log.SetFlags(0)
	in, out := make(chan string), make(chan string)

	go readThem(in, out)

	strs := []string{"a", "b", "c", "d", "e", "f"}
	for _, s := range strs {
		in <- s
	}
	close(in)
	<-out
}
