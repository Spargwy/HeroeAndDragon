package main

import "testing"

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
	crossbow = increaseWeaponDamage(crossbow)
	if crossbow.damage != wantCrossbowDamage {
		t.Fatalf("ERROR!!! This wapon damage must be equal %d, but equal %d", wantCrossbowDamage, crossbow.damage)
	}
	crossbow.damage = 7
	wantCrossbowDamage = 6
	crossbow = increaseWeaponDamage(crossbow)
	if crossbow.damage != wantCrossbowDamage {
		t.Fatalf("ERROR!!! This wapon damage must be equal %d, but equal %d", wantCrossbowDamage, crossbow.damage)
	}

	wantPanDamage := 9
	pan = increaseWeaponDamage(pan)
	if pan.damage != wantPanDamage {
		t.Fatalf("ERROR!!! This wapon damage must be equal %d, but equal %d", wantPanDamage, pan.damage)
	}

	wantSwordDamage := 30
	sword = increaseWeaponDamage(sword)
	if sword.damage != wantSwordDamage {
		t.Fatalf("ERROR!!! This wapon damage must be equal %d, but equal %d", wantSwordDamage, sword.damage)
	}

}
