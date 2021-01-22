package main

import (
	"fmt"
	"testing"
)

func TestHeroChanceToAttack(t *testing.T) {
	weaponMissChance := 0
	chanceIsGood := heroChanceToAttack(weaponMissChance, 18)
	wantChanceIsGood := true
	if chanceIsGood != wantChanceIsGood {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantChanceIsGood, chanceIsGood)
	}
	weaponMissChance = 0
	chanceIsGood = heroChanceToAttack(weaponMissChance, 19)
	wantChanceIsGood = true
	if chanceIsGood != wantChanceIsGood {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantChanceIsGood, chanceIsGood)
	}
	fmt.Print("OK\n")

}
