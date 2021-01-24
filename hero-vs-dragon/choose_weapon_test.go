package main

import (
	"fmt"
	"testing"
)

func TestUsingWeapon(t *testing.T) {
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
		ItemUsed:      9,
	}
	sword := Weapon{
		name:          "sword",
		damage:        30,
		missChance:    20,
		numberOfUsing: 1,
		ItemUsed:      0,
	}
	wantCrossbowName := crossbow.name
	wantPanName := pan.name
	wantSwordName := sword.name

	weapon, crossbow, pan := usingWeapon(crossbow, pan, sword, "1")
	if weapon.name != wantCrossbowName {
		t.Fatalf("Return weapon struct is incorrect!!! Returned %s, but most %s", weapon.name, wantCrossbowName)
	}
	weapon, crossbow, pan = usingWeapon(crossbow, pan, sword, "2")
	if weapon.name != wantPanName {
		t.Fatalf("Return weapon struct is incorrect!!! Returned %s, but most %s", weapon.name, wantPanName)
	}

	weapon, crossbow, pan = usingWeapon(crossbow, pan, sword, "3")
	if weapon.name != wantSwordName {
		t.Fatalf("Return weapon struct is incorrect!!! Returned %s, but most %s", weapon.name, wantSwordName)
	}
	fmt.Print("OK\n")

}
