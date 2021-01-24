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

func chooseWeapon() string {
	var choosedWeapon string
	fmt.Print("If you want to use weapon, you can put number of it's weapon\n")
	fmt.Printf("%s: 1\n", "crossbow")
	fmt.Printf("%s: 2\n", "pan")
	fmt.Printf("%s: 3\n", "sword")
	fmt.Printf("weapon: ")

	fmt.Scan(&choosedWeapon)

	return choosedWeapon
}

func usingWeapon(choosedWeapon string) Weapon {
	crossbow := Weapon{
		name:          "crossbow",
		damage:        30,
		minDamage:     15,
		missChance:    30,
		numberOfUsing: 5,
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
	var weapon Weapon

	switch choosedWeapon {
	case "1":
		if crossbowItemUsed < crossbow.numberOfUsing {
			weapon = crossbow
			crossbowItemUsed++
		} else {
			fmt.Printf("you have exceeded your weapon usage limit%d. Please choose another weapon\n", crossbowItemUsed)
			choosedWeapon := chooseWeapon()
			usingWeapon(choosedWeapon)
		}
	case "2":
		if panItemUsed < pan.numberOfUsing {
			weapon = pan
			panItemUsed++
		} else {
			fmt.Printf("you have exceeded your weapon usage limit%d. Please, choose another weapon\n", panItemUsed)
			choosedWeapon := chooseWeapon()
			usingWeapon(choosedWeapon)
		}
	case "3":
		weapon = sword

	default:
		fmt.Print("You put the incorrect number, please, enter again\n")
		choosedWeapon := chooseWeapon()
		usingWeapon(choosedWeapon)
	}
	return weapon
}
