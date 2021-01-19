package main

import "fmt"

func usingWeapon(usedTimes int, weaponName string, weaponDamage int) Weapon {
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
	standartSword.damage = 10
	standartSword.missChacnce = 0
	standartSword.numberOfUses = 1
	standartSword.usedTimes = 0

	var weapon Weapon
	switch weaponName {
	case "pan":
		if usedTimes < pan.numberOfUses {
			weapon = pan
			weapon.damage = weaponDamage

		} else {
			fmt.Printf("you have reached the limit for this weapon(%d), used default weapon \n(to disable the limits on weapons, you can activate the service 'endless ammunition'for 1,99$)\n",
				weapon.usedTimes)
			weapon = standartSword
		}
	case "crossbow":
		if usedTimes < crossbow.numberOfUses {
			weapon = crossbow
			weapon.damage = weaponDamage

		} else {
			fmt.Printf("you have reached the limit for this weapon(%d), used default weapon \n(to disable the limits on weapons, you can activate the service 'endless ammunition'for 1,99$)\n",
				weapon.usedTimes)
			weapon = standartSword
			weapon.damage = weaponDamage
		}
	default:
		weapon = standartSword
		weapon.damage = weaponDamage
		fmt.Printf("Used default weapon\n")
		fmt.Printf("%d hhhhhhhhhhhh\n\n\n", weapon.damage)

	}

	return weapon
}
