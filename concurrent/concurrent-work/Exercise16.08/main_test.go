package main

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func Test_main(t *testing.T) {
	var s bytes.Buffer
	log.SetOutput(&s)
	log.SetFlags(0)

	main()

	output := s.String()
	ln := len(output)
	res := output[ln-5 : ln-1]

	if res != "5050" {
		t.Errorf("Expected 5050 but received %s", res)
	}
}

func Test_sum(t *testing.T) {
	for i := 4; i < 20; i++ {
		res := sum(i, 1, 100)
		if res != 5050 {
			t.Errorf("We were expecting 5050 with %d workers but we received %d", i, res)
		}
	}
}

func Benchmark_Sum(b *testing.B) {
	for i := 4; i < 20; i++ {
		b.Run(fmt.Sprintf("=> %d workers: \n", i), func(b *testing.B) {
			for c := 0; c < b.N; c++ {
				res := sum(i, 1, 100)
				if res != 5050 {
					b.Errorf("We were expecting 5050 with %d workers but we received %d", i, res)
				}
			}
		})
	}
}

/**

Run specific benchmark:

go test -benchmem -run=^$ -bench ^(Benchmark_Sum)$

*/
