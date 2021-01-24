package main

import (
	"fmt"
	"testing"
)

func TestChance(t *testing.T) {
	chance := chanceToAttack()
	if chance > -1 && chance < 101 {
		fmt.Print("OK")
	} else {
		t.Fail()
	}

}
