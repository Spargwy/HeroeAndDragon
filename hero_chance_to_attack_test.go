package main

import "testing"

func TestHeroChanceToAttack(t *testing.T) {
	weaponMissChance := 100
	chanceIsGood := heroChanceToAttack(weaponMissChance)
	wantChanceIsGood := false
	if chanceIsGood != wantChanceIsGood {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantChanceIsGood, chanceIsGood)
	}
	weaponMissChance = 0
	chanceIsGood = heroChanceToAttack(weaponMissChance)
	wantChanceIsGood = true
	if chanceIsGood != wantChanceIsGood {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantChanceIsGood, chanceIsGood)
	}
}
