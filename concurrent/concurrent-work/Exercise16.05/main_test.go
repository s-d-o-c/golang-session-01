package main

import (
	"bytes"
	"log"
	"testing"
)

func Test_main(t *testing.T) {
	var s bytes.Buffer
	log.SetOutput(&s)
	log.SetFlags(0)

	main()

	if s.String() != "Thanks for Hello Hari\nHello Muhammad\n" {
		t.Error(s.String())
	}
}
