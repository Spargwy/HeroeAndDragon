package main

import "fmt"

type Weapon struct {
	name              string
	damage            int
	minDamage         int
	missChance        int
	numberOfUsing     int
	ItemUsed          int
	minChanceToAttack int
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
	if choosedWeapon == "\n" || choosedWeapon == "" {
		fmt.Printf("Please, enter the correct value\n")
	}
	return choosedWeapon
}

func usingWeapon(crossbow, pan, sword Weapon, choosedWeapon string) (Weapon, Weapon, Weapon) {
	var weapon Weapon
	switch choosedWeapon {
	case "1":
		if crossbow.ItemUsed < crossbow.numberOfUsing {
			weapon = crossbow
			crossbow.ItemUsed++

		} else {
			fmt.Printf("you have exceeded your weapon usage limit%d. Please choose another weapon\n", crossbow.ItemUsed)
			choosedWeapon := chooseWeapon()
			usingWeapon(crossbow, pan, sword, choosedWeapon)
		}
	case "2":
		if pan.ItemUsed < pan.numberOfUsing {
			weapon = pan
			pan.ItemUsed++
		} else {
			fmt.Printf("you have exceeded your weapon usage limit%d. Please, choose another weapon\n", pan.ItemUsed)
			choosedWeapon := chooseWeapon()
			usingWeapon(crossbow, pan, sword, choosedWeapon)
		}
	case "3":
		weapon = sword

	default:
		fmt.Print("You put the incorrect number, please, enter again\n")
		choosedWeapon := chooseWeapon()
		usingWeapon(crossbow, pan, sword, choosedWeapon)
	}
	return weapon, crossbow, pan
}
