package main

import "fmt"

type Weapon struct {
	name          string
	damage        int
	minDamage     int
	missChance    int
	numberOfUsing int
	ItemUsed      int
}

//принимать параметры сложности
func usingWeapon() Weapon {
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

	var WeaponChoose string
	var weapon Weapon
	fmt.Print("If you want to use weapon, you can put number of it's weapon\n")
	fmt.Printf("%s: 1\n", crossbow.name)
	fmt.Printf("%s: 2\n", pan.name)
	fmt.Printf("weapon: ")

	fmt.Scan(&WeaponChoose)

	switch WeaponChoose {
	case "1":
		if crossbowItemUsed < crossbow.numberOfUsing {
			weapon = crossbow
			crossbowItemUsed++
		} else {
			fmt.Printf("you have exceeded your weapon usage limit. Standard sword used, %d\n", crossbowItemUsed)
			weapon = sword
		}
	case "2":
		if panItemUsed < pan.numberOfUsing {
			weapon = pan
			panItemUsed++
		} else {
			fmt.Printf("you have exceeded your weapon usage limit. Standard sword used, %d\n", panItemUsed)
			weapon = sword
		}
	default:
		fmt.Print("You put the incorrect number, used default weapon - sword\n")
		weapon = sword
	}
	return weapon
}
