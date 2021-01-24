package main

import (
	"testing"
)

func TestDragonChanceToAttack(t *testing.T) {
	dragonMissChance := 100
	chanceIsGood := dragonChanceToAttack(dragonMissChance)
	wantChanceIsGood := false
	if chanceIsGood != wantChanceIsGood {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantChanceIsGood, chanceIsGood)
	}
	dragonMissChance = 0
	chanceIsGood = dragonChanceToAttack(dragonMissChance)
	wantChanceIsGood = true
	if chanceIsGood != wantChanceIsGood {
		t.Fatalf("ERROR, chance is good must be equal %v, but equal %v", wantChanceIsGood, chanceIsGood)
	}
}
