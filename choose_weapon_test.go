package main

import "testing"

func TestUsingWeapon(t *testing.T) {
	crossbow := Weapon{}
	crossbow.name = "crossbow"
	crossbow.damage = 30
	crossbow.minDamage = 15
	crossbow.missChance = 20
	crossbow.numberOfUsing = 100
	crossbow.ItemUsed = 0

	pan := Weapon{}
	pan.name = "pan"
	pan.damage = 10
	pan.missChance = 0
	pan.numberOfUsing = 10
	pan.ItemUsed = 0

	sword := Weapon{}
	sword.name = "sword"
	sword.damage = 30
	sword.missChance = 20
	sword.numberOfUsing = 1
	sword.ItemUsed = 0

	wantCrossbowName := crossbow.name
	wantPanName := pan.name
	wantSwordName := sword.name

	weapon := usingWeapon("1")
	if weapon.name != wantCrossbowName {
		t.Fatalf("Return weapon struct is incorrect!!! Returned %s, but most %s", weapon.name, wantCrossbowName)
	}
	weapon = usingWeapon("2")
	if weapon.name != wantPanName {
		t.Fatalf("Return weapon struct is incorrect!!! Returned %s, but most %s", weapon.name, wantPanName)
	}

	weapon = usingWeapon("22424")
	if weapon.name != wantSwordName {
		t.Fatalf("Return weapon struct is incorrect!!! Returned %s, but most %s", weapon.name, wantSwordName)
	}
}
