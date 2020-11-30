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

	output := s.String()
	ln := len(output)
	res := output[ln-5 : ln-1]

	if res != "5050" {
		t.Errorf("Expected 5050 but received %s", res)
	}
}

/**
sometimes:

--- FAIL: Test_main (0.00s)
    main_test.go:21: Expected 5050 but received 5025
FAIL

--- FAIL: Test_main (0.00s)
    main_test.go:21: Expected 5050 but received 5075
FAIL

--- FAIL: Test_main (0.00s)
    main_test.go:21: Expected 5050 but received 5100
FAIL

**/
