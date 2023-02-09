package main

import (
	"fmt"
	"testing"
	"time"
)

func TestWorking(t *testing.T) {
	start := time.Now()

	Working()

	elapsed := time.Since(start)
	t.Logf("Time taken: %s", elapsed)

	if elapsed > (2 * time.Second) {
		t.Errorf("Not running under 2 sec")
	} else {
		fmt.Println("Ran under 2 secondss")
	}
}
