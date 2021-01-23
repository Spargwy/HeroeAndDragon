package main

import (
	"fmt"
	"testing"
)

func TestHeroChanceToAttack(t *testing.T) {
	weaponMissChance := 0
	successfulAttack := heroChanceToAttack(weaponMissChance, 18)
	wantsuccessfullyAttack := true
	if successfulAttack != wantsuccessfullyAttack {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantsuccessfullyAttack, successfulAttack)
	}
	weaponMissChance = 0
	successfulAttack = heroChanceToAttack(weaponMissChance, 19)
	wantsuccessfullyAttack = true
	if successfulAttack != wantsuccessfullyAttack {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantsuccessfullyAttack, successfulAttack)
	}
	fmt.Print("OK\n")

}
