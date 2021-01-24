package main

import (
	"fmt"
	"testing"
)

func TestIncreaseWeapondamage(t *testing.T) {
	crossbow := Weapon{
		name:          "crossbow",
		damage:        30,
		minDamage:     15,
		missChance:    20,
		numberOfUsing: 100,
		ItemUsed:      0,
	}
	pan := Weapon{
		name:          "pan",
		damage:        10,
		missChance:    0,
		numberOfUsing: 10,
		ItemUsed:      0,
	}
	sword := Weapon{
		name:          "sword",
		damage:        30,
		missChance:    20,
		numberOfUsing: 1,
		ItemUsed:      0,
	}

	wantCrossbowDamage := 27
	crossbow, _, _ = increaseWeaponDamage(crossbow, crossbow, pan, sword)
	if crossbow.damage != wantCrossbowDamage {
		t.Fatalf("ERROR!!! This wapon damage must be equal %d, but equal %d", wantCrossbowDamage, crossbow.damage)
	}

	crossbow.damage = 7
	wantCrossbowDamage = 6
	crossbow, _, _ = increaseWeaponDamage(crossbow, crossbow, pan, sword)
	if crossbow.damage != wantCrossbowDamage {
		t.Fatalf("ERROR!!! This wapon damage must be equal %d, but equal %d", wantCrossbowDamage, crossbow.damage)
	}

	wantPanDamage := 9
	fmt.Println(pan)
	_, pan, _ = increaseWeaponDamage(pan, crossbow, pan, sword)
	if pan.damage != wantPanDamage {
		t.Fatalf("ERROR!!! This wapon damage must be equal %d, but equal %d", wantPanDamage, pan.damage)
	}

	wantSwordDamage := 30
	_, _, sword = increaseWeaponDamage(sword, crossbow, pan, sword)
	if sword.damage != wantSwordDamage {
		t.Fatalf("ERROR!!! This wapon damage must be equal %d, but equal %d", wantSwordDamage, sword.damage)
	}

}
