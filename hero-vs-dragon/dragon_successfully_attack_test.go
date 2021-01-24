package main

import (
	"fmt"
	"testing"
)

func TestDragonSuccessfullyAttack(t *testing.T) {
	dragonMissChance := 40
	successfulAttack := dragonSuccessfullyAttack(dragonMissChance, 20)
	wantSuccessfullyAttack := true
	if successfulAttack != wantSuccessfullyAttack {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantSuccessfullyAttack, successfulAttack)
	}
	dragonMissChance = 40
	successfulAttack = dragonSuccessfullyAttack(dragonMissChance, 61)
	wantSuccessfullyAttack = false
	if successfulAttack != wantSuccessfullyAttack {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantSuccessfullyAttack, successfulAttack)
	}
	fmt.Print("OK\n")

}
