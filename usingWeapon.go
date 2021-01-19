package main

import "fmt"

func usingWeapon(usedTimes int, weaponName string) Weapon {
	pan := Weapon{}
	pan.name = "pan"
	pan.damage = 10
	pan.missChacnce = 0
	pan.numberOfUses = 2
	pan.usedTimes = 0

	crossbow := Weapon{}
	crossbow.name = "crossbow"
	crossbow.damage = 30
	crossbow.missChacnce = 10
	crossbow.numberOfUses = 2
	crossbow.usedTimes = 0

	standartSword := Weapon{}
	standartSword.name = "standart"
	standartSword.damage = 20
	standartSword.missChacnce = 20
	standartSword.numberOfUses = 1
	standartSword.usedTimes = 0

	var weapon Weapon
	switch weaponName {
	case "pan":
		if usedTimes < pan.numberOfUses {
			weapon = pan

		} else {
			fmt.Printf("you have reached the limit for this weapon(%d), used default weapon \n(to disable the limits on weapons, you can activate the service 'endless ammunition'for 1,99$)\n",
				weapon.usedTimes)
			weapon = standartSword
		}
	case "crossbow":
		if usedTimes < crossbow.numberOfUses {
			weapon = crossbow
		} else {
			fmt.Printf("you have reached the limit for this weapon(%d), used default weapon \n(to disable the limits on weapons, you can activate the service 'endless ammunition'for 1,99$)\n",
				weapon.usedTimes)
			weapon = standartSword
		}
	default:
		weapon = standartSword
		fmt.Printf("Used default weapon\n")
	}

	return weapon
}
