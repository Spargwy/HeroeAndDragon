package main

import (
	"fmt"
	"testing"
)

func TestrandomSeed(t *testing.T) {
	seed := randomSeed()
	if seed > -1 && seed < 101 {
		fmt.Print("OK")
	} else {
		t.Fail()
	}

}
