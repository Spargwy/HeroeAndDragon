package main

import (
	"fmt"
	"testing"
)

func TestDragonChanceToAttack(t *testing.T) {
	dragonMissChance := 40
	successfulAttack := dragonSuccessfullyAttack(dragonMissChance, 18)
	wantChanceIsGood := true
	if successfulAttack != wantChanceIsGood {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantChanceIsGood, successfulAttack)
	}
	dragonMissChance = 40
	successfulAttack = dragonSuccessfullyAttack(dragonMissChance, 19)
	wantChanceIsGood = false
	if successfulAttack != wantChanceIsGood {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantChanceIsGood, successfulAttack)
	}
	fmt.Print("OK\n")

}
