package main

import (
	"fmt"
)

func heroSuccessfullyAttack(weaponMissChance int, chance int) bool {
	var successfulAttack bool
	if chance > 100-weaponMissChance {
		fmt.Printf("You missed with chance %d!\n", chance)
	} else {
		successfulAttack = true
	}

	return successfulAttack
}

func heroAttack(hero Hero, dragon Dragon, crossbow, pan, sword, weapon Weapon, successfulAttack bool) (int, Weapon, Weapon, Weapon, Weapon) {
	damageReduction := heroTiredness(hero)
	if weapon.minDamage > weapon.damage-damageReduction {
		hero.damage = weapon.minDamage
	} else {
		hero.damage = weapon.damage - damageReduction
	}
	if weapon.damage > 0 {
		if successfulAttack == true {
			dragon.health = dragon.health - hero.damage

			crossbow, pan, sword = increaseWeaponDamage(weapon, crossbow, pan, sword)
			/*if weapon.name == "crossbow" {
				crossbow = increaseWeaponDamage(crossbow)
			} else if weapon.name == "pan" {
				pan = increaseWeaponDamage(pan)
			} else if weapon.name == "sword" {
				sword = increaseWeaponDamage(sword)
			}*/

		}
	} else if weapon.damage <= 0 {
		fmt.Printf("WEAPON IS BROKEN %s \n", weapon.name)
		choosedWeapon := chooseWeapon()
		usingWeapon, _, _ := usingWeapon(crossbow, pan, sword, choosedWeapon)
		heroAttack(hero, dragon, crossbow, pan, sword, usingWeapon, successfulAttack)
	}

	return dragon.health, weapon, crossbow, pan, sword
}
