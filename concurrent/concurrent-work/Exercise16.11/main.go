package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func countNumbers(c context.Context, r chan int) {
	v := 0
	for {
		select {
		case <-c.Done():
			fmt.Println("countNumbers: r<-v")
			r <- v
			break
		default:
			fmt.Println("countNumbers: time.Sleep(time.Millisecond * 100)")
			time.Sleep(time.Millisecond * 100)
			v++
		}
	}
}

func main() {
	r := make(chan int)
	c := context.TODO()
	cl, stop := context.WithCancel(c)
	go countNumbers(cl, r)

	go func() {
		fmt.Println(">>> go func(): time.Sleep(time.Millisecond * 100 * 3)")
		time.Sleep(time.Millisecond * 100 * 3)
		stop()
	}()

	v := <-r

	log.Println(v)
}
